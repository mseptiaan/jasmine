generate-proto:
	protoc -I ./pkg/proto --go_out=./internal/pb --go_opt=paths=source_relative --go-grpc_out=./internal/pb --go-grpc_opt=paths=source_relative --grpc-gateway_out=./internal/pb --grpc-gateway_opt=paths=source_relative ./pkg/proto/*.proto
	make generate-proto-tag
generate-proto-tag:
	protoc-go-inject-tag -input="./internal/pb/*.pb.go"
tidy:
	go mod tidy
install:
	make generate-proto
	make generate-proto-tag
	make tidy
