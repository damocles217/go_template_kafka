FROM golang:latest AS build-stage

WORKDIR /app/src
COPY . .

RUN go mod download

RUN go build -o /app/microservice /app/src/


FROM alpine:latest AS build-release

WORKDIR /

COPY --from=build-stage /app/microservice /usr/local/bin/microservice

EXPOSE 5000
EXPOSE 9092
EXPOSE 5432

CMD ["microservice"]