compile:
	@protoc api/v1/*.proto --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --proto_path=.

test:
	go test -race ./...

start_server:
	go run cmd/server/main.go --port 8080 --dir /var/prolog