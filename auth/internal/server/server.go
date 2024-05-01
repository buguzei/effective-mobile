package server

import (
	pb "github.com/buguzei/effective-mobile/pkg/protos/gen/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
)

func Run(auth pb.AuthServer) error {
	listener, err := net.Listen("tcp", os.Getenv("PORT"))
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterAuthServer(s, auth)

	if err = s.Serve(listener); err != nil {
		return err
	}

	return nil
}
