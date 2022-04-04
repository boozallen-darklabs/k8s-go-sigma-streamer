#!/bin/bash
declare test_job="deploy/resources/kafka_test.yaml"
kubectl apply -f ${test_job}
declare test_pod="$(kubectl get pods -o name | grep integration-test)"
kubectl wait --timeout=5s --for=condition=Ready ${test_pod}
kubectl logs -f ${test_pod} | grep -v "INFO" 2&> /dev/null
kubectl delete -f ${test_job}
