package calculate

import (
	"context"
)

type Service interface {
	Calculate(ctx context.Context, in CalculateIn) ResultAck
}

type CalculateIn struct {
	A    int    `json:"a"`
	B    int    `json:"b"`
	Type string `json:"type"`
}

type ResultAck struct {
	Res int `json:"res"`
}

type baseServer struct{}

func NewService() Service {
	return &baseServer{}
}

func (s baseServer) Calculate(ctx context.Context, in CalculateIn) ResultAck {
	if in.Type == "plus" {
		return ResultAck{Res: in.A + in.B}
	} else if in.Type == "minus" {
		return ResultAck{Res: in.A - in.B}
	} else if in.Type == "times" {
		return ResultAck{Res: in.A * in.B}
	} else if in.Type == "devide" {
		return ResultAck{Res: in.A / in.B}
	}
	return ResultAck{}
}
