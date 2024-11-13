module service-b

go 1.22.7

toolchain go1.22.9

replace service-b/proto => ./proto

require (
	google.golang.org/grpc v1.68.0
	google.golang.org/protobuf v1.35.1
)

require (
	golang.org/x/net v0.29.0 // indirect
	golang.org/x/sys v0.25.0 // indirect
	golang.org/x/text v0.18.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240903143218-8af14fe29dc1 // indirect
)