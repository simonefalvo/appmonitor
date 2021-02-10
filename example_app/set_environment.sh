#!/bin/bash

export PROMETHEUS_HOSTNAME=$(minikube ip)
export PROMETHEUS_PORT=30007
export APPLICATION_NAME=sequence-orchestrator.openfaas-fn
export QUERY_PERIOD=20
