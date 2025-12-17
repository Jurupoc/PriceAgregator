.PHONY: proto
proto:
	@echo "Generating gRPC code..."
	@mkdir -p proto/gen/price
	@protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		proto/price.proto

