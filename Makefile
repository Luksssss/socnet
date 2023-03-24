.PHONY: generate
generate:
	@ protoc -I . \
 		  --go_out ./internal/pb/ --go_opt paths=source_relative \
 		  --go-grpc_out ./internal/pb/ --go-grpc_opt paths=source_relative \
 		  --grpc-gateway_out ./internal/pb/ \
          --grpc-gateway_opt logtostderr=true \
          --grpc-gateway_opt paths=source_relative \
          ./api/socnet/soc-net.proto

gen-csv:
	@ go run ./scripts/users_from_csv.go

loads-get-users:
	@ vegeta attack -targets=./load-testing/target-get_users.list -duration=20s -rate=10 -connections=100| vegeta report
