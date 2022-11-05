FROM golang:1.18 AS build

WORKDIR /go/src/github.com/MitoVeli/math_grpc_client
COPY . .

RUN go build -o /go/bin/math_grpc_client

FROM gcr.io/distroless/base-debian10

COPY --from=build /go/bin/math_grpc_client /go/bin/math_grpc_client

ENTRYPOINT ["/go/bin/math_grpc_client"]

