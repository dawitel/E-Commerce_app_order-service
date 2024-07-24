# A gRPC-API-gateway for orderservice in a ecommerce website
<p align="center    ">
    <img src="/public/gateway-proxying.webp">
</p>

The intention is that this gRPC service will be consumed by a REST-API-gateway to perform the 
actions of creating, tracking, updating, fetching, and deleting the order and transporting it to the 
payment service.

# How to run the service locally
## using Docker
```bash
git clone "github.com/dawitel/grpc-go.git"
cd grpc-go

make protoc
docker-compose up -d
```
OR 

```bash
git clone "github.com/dawitel/grpc-go.git"
cd grpc-go

make protoc
make server
make client
```

# Folder Structure
grpc-go
├── src
│   └── server
│       └── main.go
│   └── client
│       └── main.go
├── go.mod
├── go.sum
├── internal
│   ├── db.go
│   └── orderservice.go
├── test-data
│   ├── data.json
│   └── postman_coll.json
├── Makefile
├── test.main.go
├── docler-compose.yaml
├── DockerFile
├── proto
│   ├── google
│   │   └── api
│   │       |── date.proto
|   |       |── http.proto
|   |       └── annotations.proto
│   ├── orders
│   │   └── order.proto
│   └── product
│       └── product.proto
└── protogen
    └── . . .