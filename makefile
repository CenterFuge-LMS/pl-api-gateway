proto:
	protoc -I . --go_out . --go_opt=paths=source_relative --go-grpc_out . --go-grpc_opt=paths=source_relative pkg/**/pb/*.proto


server:
	go run cmd/api-server/main.go