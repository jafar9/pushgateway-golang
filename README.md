# prometheus-pushgateway-golang
while running batch jobs we can able to see the metrics like accuracy, loss ..etc, once job got completed, it's not possible to see those metrics. if we want to see the metrics of completed job, we need to store in the time series database like prometheus. prometheus doesn't scrape the metrics from job. first we have to send those metrics to pushgateway, and then prometheus will scare it from pushgateway.

