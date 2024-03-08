# Let's Go

> In this chapter, we built a simple JSON/HTTP commit log service that accepts and responds with JSON and stores the records in those requests to an in- memory log. Next, we’ll use protocol buffers to manage our API types, generate custom code, and prepare to write a service with gRPC—an open source, high-performance remote procedure call framework that’s great for building distributed services.

## Test Your API

Start the server:

```shell
go run cmd/server/main.go
```

Add records:

```shell
curl -X POST localhost:8080 -d '{"record": { "value": "TGV0J3MgR28gIzEK" }}'

curl -X POST localhost:8080 -d '{"record": { "value": "TGV0J3MgR28gIzIK" }}'

curl -X POST localhost:8080 -d '{"record": { "value": "TGV0J3MgR28gIzMK" }}'
```

Retrieve records:

```shell
curl -X GET localhost:8080 -d '{"offset": 0}'

curl -X GET localhost:8080 -d '{"offset": 1}'

curl -X GET localhost:8080 -d '{"offset": 2}'
```
