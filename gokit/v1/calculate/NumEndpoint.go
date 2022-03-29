package calculate

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

func MakeCalculateEndPoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		in := request.(CalculateIn)
		ack := s.Calculate(ctx, in)
		return ack, nil
	}
}

type EndPointServer struct {
	CalculateEndPoint endpoint.Endpoint
}

func NewEndPointServer(svc Service) EndPointServer {
	CalculateEndPoint := MakeCalculateEndPoint(svc)
	return EndPointServer{CalculateEndPoint: CalculateEndPoint}
}
