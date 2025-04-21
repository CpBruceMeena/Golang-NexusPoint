module github.com/CpBruceMeena/golang-nexuspoint/golang-grpc-app

go 1.22.0

toolchain go1.23.1

require (
	github.com/CpBruceMeena/golang-nexuspoint/proto v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.64.0
	google.golang.org/protobuf v1.33.0
)

replace github.com/CpBruceMeena/golang-nexuspoint/proto => ../proto

require (
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250115164207-1a7da9e5054f // indirect
)
