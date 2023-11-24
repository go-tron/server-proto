package pb

import (
	socketpb "github.com/go-tron/socket/pb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func NewAny(src proto.Message) *anypb.Any {
	a, _ := anypb.New(src)
	return a
}

func NewMessage(clientId string, cmd uint32, message proto.Message) *socketpb.Message {
	content, _ := anypb.New(message)
	return &socketpb.Message{
		Body: &socketpb.MessageBody{
			Cmd:     cmd,
			Content: content,
		},
		ClientId: clientId,
	}
}

func NewMessageBytes(clientId string, cmd uint32, message proto.Message) ([]byte, error) {
	msg := NewMessage(clientId, cmd, message)
	return proto.Marshal(msg)
}
