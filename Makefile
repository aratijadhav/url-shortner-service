IMAGE_NAME=aratijadhav/url-shortner
CONTAINER_NAME=url-shortner-container
TAG=v0.1

all: build

build: 
	go build .

docker-build: 
	docker build -t $(IMAGE_NAME):$(TAG) .

docker-run: docker-build
	docker run -t -d --name $(CONTAINER_NAME) -p 8080:8080 $(IMAGE_NAME)

docker-push: docker-build
	docker image push $(IMAGE_NAME):$(TAG)

docker-stop:
	docker stop $(CONTAINER_NAME) && docker rm $(CONTAINER_NAME)

