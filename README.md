# math_grpc_client
Math operations grpc client; gets operation sign and variables from the user as flags, sends them to [math_grpc_server](https://github.com/MitoVeli/math_grpc_server) and receives the result of the operation.

## Build
Application can be build by the following command from the main application directory;

-   `go build ./cmd/math_grpc_client`
## Run
Application can be run in two ways, from the main application level:

1) Run the executable:

-   `./math_grpc_client -firstNum=9.0 -secondNum=3.2 -operationSign="/"`

    As mentioned, application expects two variables and an operation sign from the user, they can be defined as in the above example.
    Default value is 0 if the firstNum or secondNum is not provided while calling the executable. There is no default operation sign, so it is necessary to define the operationSign flag.

2) Run by Docker:

-   `docker-compose up --build`

    In case the application is run by docker, necessary flags should be provided from the Dockerfile ENTRYPOINT as in the example below;

-   `ENTRYPOINT ["/cmd/math_grpc_client", "--firstNum=10", "--secondNum=20", "--operationSign=+"]`

## Test
Tests can be run by the following command;
-   `go test ./... -v`

### Comments
Http server in main.go is also added just for demonstration purposes. In case the application is run by docker, http server keeps the application up and running.