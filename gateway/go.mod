module msa_big_tech/gateway

go 1.25.0

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.27.2
	google.golang.org/grpc v1.75.1
	msa_big_tech/auth v0.1.0
)

replace msa_big_tech/auth => ../auth

require (
	golang.org/x/net v0.41.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.28.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20250818200422-3122310a409c // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250818200422-3122310a409c // indirect
	google.golang.org/protobuf v1.36.9 // indirect
)
