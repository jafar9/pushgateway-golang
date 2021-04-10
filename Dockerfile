FROM golang:1.14

ADD pushgateway /go/pushgateway

WORKDIR /go/pushgateway

RUN go install .

ENTRYPOINT /go/bin/pushgateway

RUN rm -rf /go/pushgateway

EXPOSE 9401

