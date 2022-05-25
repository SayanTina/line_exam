# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.sum ./
COPY go.mod ./

RUN go mod download

COPY . ./

RUN go build -o /docker-backend-go

CMD ["/docker-backend-go"]

#run docker command
#docker run --publish 8080:8080 --rm docker-backend-go