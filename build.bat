cd ./proto
protoc --proto_path=./ --proto_path=../  --go_out=../../ *.proto
protoc --proto_path=./ --proto_path=../  --go-grpc_out=../../ tableMonitor.proto
protoc --proto_path=./ --proto_path=../  --go-grpc_out=../../ tableClusterMonitor.proto
protoc --proto_path=./ --proto_path=../  --go-grpc_out=../../ table.proto
protoc --proto_path=./ --proto_path=../  --go-grpc_out=../../ control.proto

cd ../pb/control
protoc-go-inject-tag -input="./control.pb.go"