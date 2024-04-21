module github.com/irisco88/protos

go 1.20

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.59.0

require (
	github.com/golang/protobuf v1.5.3
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	golang.org/x/net v0.19.0
	google.golang.org/grpc v1.57.0
	google.golang.org/protobuf v1.31.0
)

require (
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto v0.0.0-20230822172742-b8732ec3820d // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20230822172742-b8732ec3820d // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230822172742-b8732ec3820d // indirect
)
