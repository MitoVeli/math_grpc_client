# syntax=docker/dockerfile:1
# specify the base image to  be used for the application, alpine or ubuntu
FROM golang:1.17-alpine
# create a working directory inside the image
WORKDIR /app
# copy Go modules and dependencies to image
COPY go.mod go.sum ./
# download Go modules and dependencies
RUN go mod download
# copy directory files i.e all files ending with .go
COPY . .
# compile application
RUN go build -o /cmd/math_grpc_client ./cmd/math_grpc_client
# tells Docker that the container listens on specified network ports at runtime
EXPOSE 8008

ENV APP_ENV=test
# command to be used to execute when the image is used to start a container
ENTRYPOINT ["/cmd/math_grpc_client", "--firstNum=10", "--secondNum=20", "--operationSign=-"]