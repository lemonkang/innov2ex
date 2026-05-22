.PHONY: proto test

proto:
	protoc --go_out=. --go_opt=module=demo04 --go-grpc_out=. --go-grpc_opt=module=demo04 protobuf/oauth.proto

test:
	go test ./...