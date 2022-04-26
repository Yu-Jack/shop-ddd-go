build:
	docker build services/order/ -t jk82421/order:v1
	docker build services/consumer/ -t jk82421/consumer:v1

push:
	docker push jk82421/order:v1
	docker push jk82421/consumer:v1

clear:
	docker rmi jk82421/order:v1
	docker rmi jk82421/consumer:v1

restart-service:
	kubectl delete -f k8s.yaml
	kubectl delete -f istio.yaml
	istioctl kube-inject -f k8s.yaml | kubectl apply -f -
	kubectl apply -f istio.yaml

up:
	istioctl kube-inject -f k8s.yaml | kubectl apply -f -
	istioctl kube-inject -f kafka.yaml | kubectl apply -f -
	kubectl apply -f istio.yaml

down: 
	kubectl delete -f k8s.yaml
	kubectl delete -f kafka.yaml
	kubectl delete -f istio.yaml