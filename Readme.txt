URL_SHORTNER

docker build -f build/package/Dockerfile -t url-shortner:latest .

docker run url-shortner:latest


what is a contaienr


make docker-build

make docker-push

kubectl apply -f deployments/manifests/deployment-dev.yml

kubectl rollout restart deployment/golang-backend-dev -n development 
// only change when only code is changed

kubectl port-forward deployment/golang-backend-dev 8080:8080 -n development




----- to check for deployments ------

kubectl get pods -n development

kubectl logs deployment/golang-backend-dev -n development
