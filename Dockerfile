FROM golang:latest
LABEL authors="Arati Jadhav aratijadhav@gmail.com"


WORKDIR /app

COPY shortnersvc ./shortnersvc

WORKDIR /app/shortnersvc
RUN ls /app/shortnersvc

ENTRYPOINT ["go", "run", "main.go"]