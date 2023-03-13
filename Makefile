.PHONY: generate
generate:
	@ protoc -I ./api/ --go_out=plugins=grpc:./internal/pb ./api/soc-net.proto