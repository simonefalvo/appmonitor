package main

import (
	"context"
	"fmt"
	"os"
	"time"
    "strconv"
    "strings"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

func main() {

    appname , ok := os.LookupEnv("APPLICATION_NAME")
    if !ok {
        fmt.Printf("Application name not set\n")
    }
    prometheus, ok := os.LookupEnv("PROMETHEUS_HOSTNAME")
    if !ok {
        fmt.Printf("Prometheus hostname not set\n")
    }
    port, ok := os.LookupEnv("PROMETHEUS_PORT")
    if !ok {
        fmt.Printf("Prometheus port not set\n")
    }
    period, ok := os.LookupEnv("QUERY_PERIOD")
    if !ok {
        fmt.Printf("Query period not set\n")
    }
    period_seconds, err := strconv.Atoi(period)
    if err != nil{
        fmt.Printf("Error converting query period seconds\n")
        os.Exit(1)
    }

	client, err := api.NewClient(api.Config{
		Address: "http://" + prometheus + ":" + port,
	})
	if err != nil {
		fmt.Printf("Error creating client: %v\n", err)
		os.Exit(1)
	}

	v1api := v1.NewAPI(client)

    var b strings.Builder
    fmt.Fprintf(&b, "rate(gateway_functions_seconds_sum{function_name='%s'}[20s]) / ", appname)
	fmt.Fprintf(&b, "rate(gateway_functions_seconds_count{function_name='%s'}[20s])", appname)
	query := b.String()
    fmt.Println("Query: " + query)

    for {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        defer cancel() //TODO: avoid stacking
        result, warnings, err := v1api.Query(ctx, query, time.Now())
        if err != nil {
            fmt.Printf("Error querying Prometheus: %v\n", err)
            os.Exit(1)
        }
        if len(warnings) > 0 {
            fmt.Printf("Warnings: %v\n", warnings)
        }
        fmt.Printf("Result:\n%v\n", result)
        time.Sleep(time.Duration(period_seconds) * time.Second)
    }
}
