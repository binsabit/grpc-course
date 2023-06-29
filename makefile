gen: 
	protoc --go_out=protogen --go-grpc_out=protogen ./proto/*

clean:
	rm protogen/*

server:
	go run ./cmd/server/main.go -port 8080
client:
	go run ./cmd/client/main.go -address 0.0.0.0:8080
test:
	go test -cover -race ./...