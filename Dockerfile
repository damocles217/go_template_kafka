FROM golang:alpine

WORKDIR /app/src
COPY . .

RUN go build -o /app/bin