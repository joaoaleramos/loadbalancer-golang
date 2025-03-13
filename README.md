# Project loadbalancer

This Project its is a simple load balancer that distributes incoming network traffic across multiple servers  using  weighted algorithm.

## Getting Started

Basically to set up services for this load balancer you need to create a configuration file `loadfig.yaml` with the following structure:

```yaml
services:
  - name: "service1"
    url: "http://localhost:8081"
    weight: 0.5
  - name: "service2"
    url: "http://localhost:8082"
    weight: 0.2
  - name: "service3"
    url: "http://localhost:8083"
    weight: 0.3
```

you can see the example in the `loadfig.example.yaml`file in the root of the project.

After that you can run the application with the command `make run` or `go run cmd/main.go` and the application will start listening on port 8080. Make sure to weight the services according to the amount of traffic you want to distribute the weight is a number that represents the amount of traffic you want to distribute to that service and all the services should have a weight greater than 0 and max value of 1.

Run build make command with tests

```bash
make all
```

Build the application

```bash
make build
```

Run the application

```bash
make run
```

Live reload the application:

```bash
make watch
```

Run the test suite:

```bash
make test
```

Clean up binary from the last build:

```bash
make clean
```
