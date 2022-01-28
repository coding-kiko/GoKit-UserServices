module github.com/fCalixto-Gb/Final-Project/HTTPService

go 1.13

replace github.com/fCalixto-Gb/Final-Project/GRPCServiceA => ../GRPCServiceA

require (
	github.com/fCalixto-Gb/Final-Project/GRPCServiceA v0.0.0-20220125185321-1247f3b24dfe
	github.com/go-kit/kit v0.12.0
	github.com/go-kit/log v0.2.0
	github.com/gorilla/mux v1.8.0
	github.com/stretchr/testify v1.7.0
	google.golang.org/grpc v1.44.0
)
