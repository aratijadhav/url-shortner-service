
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

# Sample requests


## To generate short URL

```
curl --location --request POST 'http://localhost:8080/api/v1/createurl/www.youtube.com'
curl --location --request POST 'http://localhost:8080/api/v1/createurl/www.youtube.com'

curl --location --request POST 'http://localhost:8080/api/v1/createurl/www.stackoverflow.com'
curl --location --request POST 'http://localhost:8080/api/v1/createurl/www.github.com'
curl --location --request POST 'http://localhost:8080/api/v1/createurl/www.facebook.com'
```

## To visit original URL
Above post request will generate short URL something like this: 
```
Shortned URL for www.youtube.com is: http://localhost:8080/LxcDI
```
Just visit following URL in browser:
```
http://localhost:8080/LxcDI
```


## To get mostvisited URL:
```
curl --location 'http://localhost:8080/api/v1/mostvisited'
```