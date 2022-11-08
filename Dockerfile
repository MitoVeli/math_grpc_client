# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /cmd/math_grpc_client

EXPOSE 8008

CMD [ "/math_grpc_client" ]