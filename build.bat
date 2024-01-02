cd ./proto/game
protoc --proto_path=./ --proto_path=../  --go_out=../../../ *.proto

cd ../user
protoc --proto_path=./ --proto_path=../  --go_out=../../../ *.proto

cd ../control
protoc --proto_path=./ --proto_path=../  --go_out=../../../ *.proto

cd ../table
protoc --proto_path=./ --proto_path=../  --go_out=../../../ *.proto
protoc --proto_path=./ --proto_path=../   --go-grpc_out=../../../ table.proto
protoc --proto_path=./ --proto_path=../   --go-grpc_out=../../../ tableCluster.proto

cd ../../pb/control
protoc-go-inject-tag -input="./control.pb.go"