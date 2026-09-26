IMAGE_NAME = tusand/go-hello-world
TAG = v4

.PHONY: docker-build docker-push

docker-build:
	docker build -f build/package/Dockerfile -t $(IMAGE_NAME):$(TAG) .

docker-login:
	docker login

docker-push: docker-login
	docker push $(IMAGE_NAME):$(TAG)
