package grpc

import (
	"context"
	"github.com/buguzei/effective-mobile/cars/internal/models"
	"github.com/buguzei/effective-mobile/cars/internal/usecase"
	"github.com/buguzei/effective-mobile/pkg/logging"
	pb "github.com/buguzei/effective-mobile/pkg/protos/gen/cars"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CarsHandler struct {
	l logging.Logger
	pb.UnimplementedCarsServer
	uc usecase.CarsUCI
}

func New(uc usecase.CarsUCI) *CarsHandler {
	var logger logging.Logger = logging.NewLogrus("debug")
	logger = logger.Named("usecase.CarsHandler")

	return &CarsHandler{l: logger, uc: uc}
}

func (h CarsHandler) GetCars(ctx context.Context, in *pb.GetCarsRequest) (*pb.GetCarsResponse, error) {
	const f = "CarsHandler.GetCars"

	car := models.Car{
		RegNum:  in.GetCar().GetRegNum(),
		Mark:    in.GetCar().GetMark(),
		Model:   in.GetCar().GetModel(),
		OwnerID: int(in.GetCar().GetOwnerID()),
	}

	cars, err := h.uc.GetCars(ctx, car)
	if err != nil {
		h.l.Error("error of CarsUCI.GetCars", logging.Fields{
			"error": err,
			"func":  f,
		})

		return nil, status.Error(codes.Internal, f)
	}

	protoCars := make([]*pb.Car, len(cars))
	for i, car := range cars {
		protoCars[i] = &pb.Car{}

		protoCars[i].Model = car.Model
		protoCars[i].RegNum = car.RegNum
		protoCars[i].Mark = car.Mark
		protoCars[i].OwnerID = int32(car.OwnerID)
	}

	resp := &pb.GetCarsResponse{
		Cars: protoCars,
	}

	return resp, nil
}

func (h CarsHandler) NewCar(ctx context.Context, in *pb.NewCarRequest) (*pb.NewCarResponse, error) {
	const f = "CarsHandler.NewCar"

	car := models.Car{
		RegNum:  in.GetCar().GetRegNum(),
		Mark:    in.GetCar().GetMark(),
		Model:   in.GetCar().GetModel(),
		OwnerID: int(in.GetCar().GetOwnerID()),
	}

	err := h.uc.NewCar(ctx, car)
	if err != nil {
		h.l.Error("error of CarsUCI.NewCar", logging.Fields{
			"error": err,
			"func":  f,
		})
	}

	resp := &pb.NewCarResponse{
		Status: "OK",
	}

	return resp, nil
}

func (h CarsHandler) UpdateCar(ctx context.Context, in *pb.UpdateCarRequest) (*pb.UpdateCarResponse, error) {
	const f = "CarsHandler.UpdateCar"

	filter := models.Car{
		RegNum:  in.GetCar().GetRegNum(),
		Mark:    in.GetCar().GetMark(),
		Model:   in.GetCar().GetModel(),
		OwnerID: int(in.GetCar().GetOwnerID()),
	}

	err := h.uc.UpdateCar(ctx, filter)
	if err != nil {
		h.l.Error("error of CarsUCI.UpdateCar", logging.Fields{
			"error": err,
			"func":  f,
		})
	}

	resp := &pb.UpdateCarResponse{
		Status: "OK",
	}

	return resp, nil
}

func (h CarsHandler) DeleteCar(ctx context.Context, in *pb.DeleteCarRequest) (*pb.DeleteCarResponse, error) {
	const f = "CarsHandler.DeleteCar"

	err := h.uc.DeleteCar(ctx, in.GetRegNum())
	if err != nil {
		h.l.Error("error of CarsUCI.DeleteCar", logging.Fields{
			"error": err,
			"func":  f,
		})
	}

	resp := &pb.DeleteCarResponse{
		Status: "OK",
	}

	return resp, nil
}
