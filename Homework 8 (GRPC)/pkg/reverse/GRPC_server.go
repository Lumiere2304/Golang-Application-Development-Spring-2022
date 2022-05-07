package reverse

import (
	"context"
	"gprcProgram/pkg/api"
)

type GRPCServer struct {
}

func (s *GRPCServer) Reverse(ctx context.Context, mes *api.Message) (*api.Message, error) {
	return &api.Message{
		Text: ReverseString(mes.GetText()),
	}, nil
}