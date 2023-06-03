compile:
	protoc -I. --go_out=. --go-grpc_out=. proto/*.proto
	cp -rp proto/pb server/
	cp -rp proto/pb client/
