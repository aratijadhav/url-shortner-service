FROM golang:latest

WORKDIR /app

COPY url-shortner-service ./url-shortner-service

WORKDIR /app/url-shortner-service
RUN ls /app/url-shortner-service

ENTRYPOINT ["go", "run", "main.go"]