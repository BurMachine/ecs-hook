.PHONY = gen
gen:
	export GO111MODULE=on  # Enable module mode
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	export GO_PATH=~/go
	export PATH=$PATH:/$GO_PATH/bin
	protoc --experimental_allow_proto3_optional -I api/proto --go_out=pkg/api --go-grpc_out=pkg/api api/proto/grpc.proto
