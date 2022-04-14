### Template web application for Golang

This application is a boilerplate for creating new Golang web applications. The use case for this boilerplate is to standardize Go microservices for maintainability and reusability. This template can be forked for more specialized microservice use cases.

### Setup and testing

From the go-web-template directory, run the following:

```console
go run ./cmd/web
```

Next open your web browser and navigate to the following URL: [http://localhost:8080/example-endpoint](http://localhost:8080/example-endpoint)

The webpage is intentionally blank. You should see the following in your terminal:

```console
INFO  ... Starting HTTP service `go-web-template` in `development` mode on port: 8080
INFO  ... Service `go-web-template` endpoint `ExampleEndpoint` reached.
```
