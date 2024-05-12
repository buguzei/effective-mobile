package server

import (
	pb "github.com/buguzei/effective-mobile/pkg/protos/gen/cars"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
)

func Run(cars pb.CarsServer) error {
	listener, err := net.Listen("tcp", "0.0.0.0:"+os.Getenv("PORT"))
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterCarsServer(s, cars)

	if err = s.Serve(listener); err != nil {
		return err
	}

	return nil
}
