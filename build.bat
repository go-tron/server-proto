cd ./proto
protoc --go_out=. *.proto
protoc --go-grpc_out=. table.proto
protoc --go-grpc_out=. tableCluster.proto
cd ../pb
protoc-go-inject-tag -input="./control.pb.go"