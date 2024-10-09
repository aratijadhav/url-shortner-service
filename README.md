
# URL shortner

URL Shortener is service that allows users to shorten long URLs. The service generates a unique short link for each long URL, which used to redirect users to the original link.



# Docker Commands

## Build the Docker Image

To create a Docker image for the application, run:
```
make docker-build

```
## Run the Application

To start the Docker container on port 8080, execute:

```
make docker-run

```
Once the HTTP server is up and running, you can use Postman to send requests and receive responses.

## Push to Docker Registry

To build the docker image and then push to docker registry

```
make docker-push

```
## Stop the Docker Container

To stop the running docker container.

```
make docker-stop

```

