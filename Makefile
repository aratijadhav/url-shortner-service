IMAGE_NAME=url-shortner
CONTAINER_NAME=url-shortner-container

all: build

build: 
	go build .

docker-build: 
	docker build -t $(IMAGE_NAME) .

docker-run: docker-build
	docker run -t -d --name $(CONTAINER_NAME) -p 8085:8080 $(IMAGE_NAME)

docker-stop:
	docker stop $(CONTAINER_NAME) && docker rm $(CONTAINER_NAME)

