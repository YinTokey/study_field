# shellcheck disable=SC2046
eval $(minikube docker-env)
#kubectl run egg-example-pod --image=egg_example --port=7001 imagePullPolicy=Never
#kubectl create -f egg_example_deployment.yaml
#kubectl create -f egg_example_service.yaml

# 创建公开deployment的service对象
kubectl apply -f egg_example_deployment.yaml
kubectl expose deployment egg-example --type=LoadBalancer --name=egg-example-service
curl $(minikube service egg-example-service --url)
