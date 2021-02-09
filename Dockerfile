FROM golang:1.15-alpine AS builder
WORKDIR /go/src/github.com/smvfal/appmonitor
COPY . .
RUN go build

FROM alpine
WORKDIR /root/
COPY --from=builder /go/src/github.com/smvfal/appmonitor/appmonitor .
ENV PROMETHEUS_HOSTNAME=prometheus.openfaas \
    PROMETHEUS_PORT=9090 \
    QUERY_PERIOD=10
CMD ["./appmonitor"]  
