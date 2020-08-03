module github.com/alx696/go-grpc-test

go 1.15

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.32.0

require (
	github.com/golang/protobuf v1.4.2
	google.golang.org/grpc v1.32.0
	google.golang.org/protobuf v1.25.0
)
