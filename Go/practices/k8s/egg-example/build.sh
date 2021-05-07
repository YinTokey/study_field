# shellcheck disable=SC2046
eval $(minikube docker-env)
kubectl run egg-example-pod --image=egg_example --port=7001 imagePullPolicy=Never
#kubectl create -f egg_example_deployment.yaml
#kubectl create -f egg_example_service.yaml
#curl $(minikube service egg-example-service --url)

