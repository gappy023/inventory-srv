module github.com/gappy023/inventory-srv

go 1.14

require (
	github.com/gappy023/basic v0.0.0-20200306090513-679c7c2f26dd
	github.com/golang/protobuf v1.3.2
	github.com/micro/go-micro/v2 v2.2.0
	microservice/part1/basic v0.0.0-00010101000000-000000000000 // indirect
)

replace microservice/part1/basic => github.com/gappy023/basic v0.0.0-20200306090513-679c7c2f26dd
