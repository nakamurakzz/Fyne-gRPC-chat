start-server:
	go run server/main.go

compile:
	protoc -I. --go_out=. --go-grpc_out=. proto/*.proto	