package main

import (
	"encoding/json"
	"flag"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	models "pushgateway/models"
	_ "pushgateway/models"
	"net/http"
	"regexp"
	"strings"
	"sync"
)

const (
	namespace = "nvidia"
)

type Exporter struct {
	pushgateway prometheus.Gauge
}

var mutex sync.Mutex
var MetricsData = make(map[string][]models.MetricModel)
var IsValidMetricName = regexp.MustCompile(`^[a-zA-Z_:][a-zA-Z0-9_:]*$`).MatchString

func main() {
	router := mux.NewRouter()
	var (
					listenAddress = flag.String("web.listen-address", ":9401", "Address to listen on for web interface and telemetry.")
		metricsPath   = flag.String("web.prometheus-path", "/metrics", "Path under which to expose metrics.")
	)
	flag.Parse()
	prometheus.MustRegister(NewExporter())
	router.Handle("/metrics", promhttp.Handler())
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
             <head><title>NVML Exporter</title></head>
             <body>
             <h1>NVML Exporter</h1>
             <p><a href='` + *metricsPath + `'>Metrics</a></p>
             </body>
             </html>`))
	})
	router.HandleFunc("/pushgateway", pushgateway)
	log.Println("Starting HTTP server on", *listenAddress)
	log.Println(http.ListenAndServe(*listenAddress, router))
}

func pushgateway(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if req.Body == nil {
		http.Error(w, "Request Payload not found", http.StatusNoContent)
		return
	}

	reqData := models.RequestData{}
	err := json.NewDecoder(req.Body).Decode(&reqData)
	if err != nil {
		http.Error(w, "Unable to decode the request payload", http.StatusInternalServerError)
		return
	}

	defer req.Body.Close()

	log.Printf("after decoding mlflow metric  %+v", reqData)

	mtype := reqData.GetMetricType()
	//only Gauge supported
	if mtype != "" && strings.ToLower(mtype) != "gauge" {
		http.Error(w, "only Metric Type Gauge is supported", http.StatusNotImplemented)
		return
	}

	metrics := reqData.GetMetrics()

	for _, metric := range metrics {
		key := metric.GetName()
		if IsValidMetricName(key) {
			clabels := reqData.GetCommonLabels()
			metric.Labels = append(metric.Labels, clabels...)
			if err := registerNewMetric(key, metric.Labels); err == nil {
				prepareMetricsData(key, metric)
			}
		}
	}
}

func prepareMetricsData(key string, data *models.MetricModel) {
	newMetrics := []models.MetricModel{*data}
	mutex.Lock()
	if oldData, ok := MetricsData[key]; ok {
		newMetrics = append(newMetrics, oldData...)
	}
	MetricsData[key] = newMetrics
	mutex.Unlock()
}

func registerNewMetric(key string, labels []*models.LabelModel) error {
	label_keys := []string{}
	for _, label := range labels{
		label_keys= append(label_keys, label.Key)
	}
	ml_metric := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      key,
			Help:      "pushgateway metric",
		},
		label_keys,
	)

	if err := prometheus.Register(ml_metric); err != nil {
		if _, ok := err.(prometheus.AlreadyRegisteredError); !ok {
			log.Printf("%s metric register failed, err %s\n", key, err.Error())
			return err
		}
	}
	return nil
}

func(e *Exporter) Collect(metrics chan<- prometheus.Metric) {
	scrapeMetrics(e, metrics)
}

func (e *Exporter) Describe(descs chan<- *prometheus.Desc) {
	e.pushgateway.Describe(descs)
}

func scrapeMetrics(e *Exporter, metrics chan<- prometheus.Metric) {
	mutex.Lock()
	for key, _ := range MetricsData {
	    data := MetricsData[key]
		for _, metric:= range data {
			label_keys := []string{}
			label_values := []string{}
			for _, label := range metric.Labels {
				label_keys =  append(label_keys, label.Key)
				label_values = append(label_values, label.Value)
			}
			desc := prometheus.NewDesc(metric.Name, "pushgate way metric", label_keys, nil)
			metrics <- prometheus.MustNewConstMetric(
				desc,
                prometheus.GaugeValue,
                metric.Value,
                label_values...,
            )
        }
        log.Println("dadadad %+v\n", data)
        delete(MetricsData, key)
	}
	mutex.Unlock()
}

func NewExporter() *Exporter {
	return &Exporter{
		pushgateway: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "pushgateway_metric",
				Help:      "pushgateway metric collection",
			},
		),
	}
}

