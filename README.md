# appmonitor
Simple monitoring application that makes use of Prometheus APIs to periodically check average execution time of an OpenFaaS function on a Kubernetes cluster.

## Getting Started
### Deployment
Edit [deployment](kubernetes/deployment.yml) file in order to match the OpenFaaS function you want to monitor and set your preferred query period.
```yml
        ...     
      
        env:
          - name: APPLICATION_NAME
            value: "sequence-orchestrator.openfaas-fn"
          - name: QUERY_PERIOD
            value: "20"
```
You can also specify PROMETHEUS_HOSTNAME and PROMETHEUS_PORT environment variables whether they differ from default values *prometheus.openfaas* and *9090* respectively.

Deploy the application as a kubernetes deployment as follow:
```bash
$ kubectl apply -f kubernetes/deployment.yml
```
### Running Example
The example application consists of a sequence composition of functions, *cd* into the [example_app](example_app) directory and deploy your function with the OpenFaaS CLI:
```bash
example_app$ faas-cli deploy -f stack-sequence.yml
```
Then you can watch the query result as your function gets some requests:
```bash
$ kubectl logs -f deploy/appmonitor
```
## License
Licensed under the Apache License, Version 2.0 - see the [LICENSE](LICENSE) file for details
## Acknowledgements
https://github.com/prometheus/client_golang/blob/master/api/prometheus/v1/example_test.go
