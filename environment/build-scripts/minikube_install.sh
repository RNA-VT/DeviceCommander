#!/bin/sh

minikube start --cpus 4 --memory 8192 --vm-driver=virtualbox

eval $(minikube docker-env)


# docker build -t gofire_gofire:latest .

# kubectl run discovery –-image=gofire_gofire:latest –port=8761 --image-pull-policy=Never




docker build -t gofire_gofire:latest .
kubectl run gofire --image=dtp263/gofire:v1 --image-pull-policy=Never --generator=run-pod/v1

kubectl get all | grep gofire