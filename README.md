# Distributed Tracing Example

This project contain examples of golang distributed tracing using opentracing and jaeger.

## Prerequisites

* Docker
* Golang

## How to Run This Project

Run:
```
docker-compose up
```

This command will do:
1. build server image
2. build messaging image
3. pull nats-streaming image (if you dont have one)
4. pull jaeger image (if you dont have one)

## Jaeger UI

```
http://localhost:16686/
```

## Client

### POST

Add Post

```
go run client/post/main.go
```

### GET

Retrieve Posts
```
go run client/get/main.go
```


## Test Scenario

1. docker-compose up
2. open Jaeger UI
3. run one of the client
4. check your trace in jaeger UI

## Note
You also can do request from `postman`

### Get
```
http://localhost:8080/post
```

### Post
```
http://localhost:8080/post
```

Payload:
```
{
    "title": "some text",
    "body": "some text body",
}
```
