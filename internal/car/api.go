//go:generate protoc -I=../../api --go_out=paths=source_relative,plugins=grpc:. car.proto
package car

import (
	"context"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

const BIND = "0.0.0.0:18901"

var control func(direction Direction)

func Setup(c func(direction Direction)) error {
	control = c
	server := grpc.NewServer()
	carAPI := NewCarServerAPI()
	RegisterCarRpcServer(server, carAPI)
	ln, err := net.Listen("tcp", BIND)
	if err != nil {
		return errors.Wrap(err, "start api listener error")
	}
	err = server.Serve(ln)
	if err != nil {
		logrus.Fatal(err)
	}
	return nil
}

type grpcServerAPI struct{}

func NewCarServerAPI() *grpcServerAPI {
	return &grpcServerAPI{}
}

func (c *grpcServerAPI) Command(context context.Context, request *CarControlRequest) (*CarControlResponse, error) {
	logrus.WithField("direction", request.Direction.String()).Info("internal/car: received direction")
	control(request.Direction)
	return &CarControlResponse{}, nil
}
