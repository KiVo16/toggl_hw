openapi:
	oapi-codegen -old-config-style -generate chi-server -o ./ports/http/openapi_server.gen.go -package ports ./api/openapi/api.yml
	oapi-codegen -old-config-style -generate types -o ./ports/http/openapi_types.gen.go -package ports ./api/openapi/api.yml
	
proto:
	protoc \
	--proto_path=api/protobuf api/protobuf/api.proto \
	--go_out=ports/grpc/proto --go_opt=paths=source_relative \
	--go-grpc_opt=require_unimplemented_servers=false \
	--go-grpc_out=ports/grpc/proto --go-grpc_opt=paths=source_relative
