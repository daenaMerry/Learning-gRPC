PHONY: vendor
vendor:
	go mod tidy
	go mod vendor

PHONY: run-server
run-server:
	go run server/server.go

PHONY: run-client
run-client:
	go run client/client.go

PHONY: generate_grpc
generate_grpc:
	protoc \
	--go_out=. \
	--go_opt=paths=source_relative \
	--go-grpc_out=. \
	--go-grpc_opt=paths=source_relative \
	chat/chat.proto
