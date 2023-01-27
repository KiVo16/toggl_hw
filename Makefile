openapi:
	oapi-codegen -old-config-style -generate chi-server -o ./ports/http/openapi_server.gen.go -package ports ./api/openapi/api.yml
	oapi-codegen -old-config-style -generate types -o ./ports/http/openapi_types.gen.go -package ports ./api/openapi/api.yml
	
proto:
	protoc \
	--proto_path=api/protobuf api/protobuf/api.proto \
	--go_out=ports/grpc/proto --go_opt=paths=source_relative \
	--go-grpc_opt=require_unimplemented_servers=false \
	--go-grpc_out=ports/grpc/proto --go-grpc_opt=paths=source_relative

test:
	go test ./...

build-binary:
	GOOS=linux GOARCH=arm64 go build -o ./build/build-linux-arm64 ./cmd/main.go && \
	GOOS=linux GOARCH=amd64 go build -o ./build/build-linux-amd64 ./cmd/main.go && \
	GOOS=linux GOARCH=386 go build -o ./build/build-linux-386 ./cmd/main.go && \
	GOOS=windows GOARCH=386 go build -o ./build/build-windows-386 ./cmd/main.go && \
	GOOS=windows GOARCH=amd64 go build -o ./build/build-windows-amd64 ./cmd/main.go	