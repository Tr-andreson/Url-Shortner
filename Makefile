IMAGE_NAME = tusand/go-hello-world
TAG = v1

.PHONY: docker-build docker-push

run-prod:
	APP_ENV=production go run ./cmd/server/main.go

run:
	go run ./cmd/server/main.go

docker-build:
	docker build -f build/package/Dockerfile -t $(IMAGE_NAME):$(TAG) .

docker-login:
	docker login

docker-push: docker-login
	docker push $(IMAGE_NAME):$(TAG)
