package grpc

import (
	"context"
	"github.com/buguzei/effective-mobile/auth/internal/models"
	"github.com/buguzei/effective-mobile/auth/internal/usecase"
	"github.com/buguzei/effective-mobile/pkg/logging"
	pb "github.com/buguzei/effective-mobile/pkg/protos/gen/auth"
)

type AuthHandler struct {
	l logging.Logger
	pb.UnimplementedAuthServer
	uc usecase.AuthUCI
}

func New(uc usecase.AuthUCI) *AuthHandler {
	var logger logging.Logger = logging.NewLogrus("debug")
	logger = logger.Named("auth.handlers")

	return &AuthHandler{uc: uc, l: logger}
}

func (h AuthHandler) Refresh(ctx context.Context, in *pb.RefreshRequest) (*pb.RefreshResponse, error) {
	const f = "AuthHandler.Refresh"

	pair, err := h.uc.Refresh(ctx, in.GetEmail(), in.GetRefreshToken())
	if err != nil {
		h.l.Error("error of refresh", logging.Fields{
			"error": err,
			"func":  f,
		})

		return &pb.RefreshResponse{}, err
	}

	resp := &pb.RefreshResponse{
		RefreshToken: pair.Refresh,
		AccessToken:  pair.Access,
	}

	return resp, nil
}

func (h AuthHandler) SignIn(ctx context.Context, in *pb.SignInRequest) (*pb.SignInResponse, error) {
	const f = "AuthHandler.SignIn"

	user := models.User{
		Email:    in.GetEmail(),
		Password: in.GetPassword(),
	}

	pair, err := h.uc.SignIn(ctx, user)
	if err != nil {
		h.l.Error("error of sign in", logging.Fields{
			"error": err,
			"func":  f,
		})

		return nil, err
	}

	resp := &pb.SignInResponse{
		RefreshToken: pair.Refresh,
		AccessToken:  pair.Access,
	}

	return resp, nil
}

func (h AuthHandler) SignUp(ctx context.Context, in *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	const f = "AuthHandler.SignUp"

	user := models.User{
		Name:     in.GetUsername(),
		Email:    in.GetEmail(),
		Password: in.GetPassword(),
	}

	pair, err := h.uc.SignUp(ctx, user)
	if err != nil {
		h.l.Error("error of sign up", logging.Fields{
			"error": err,
			"func":  f,
		})
		return nil, err
	}

	resp := &pb.SignUpResponse{
		RefreshToken: pair.Refresh,
		AccessToken:  pair.Access,
	}

	return resp, nil
}

//У вас есть структура Task, представляющая задачу:
//type Task struct {
//	ID       int
//	Name     string
//	Duration time.Duration // Продолжительность выполнения задачи
//}
//
//
//Необходимо написать функцию RunParallelTasks с сигнатурой:
//func RunParallelTasks(tasks []Task) {
//	ch := make(chan int, len(tasks))
//
//	for _, task := range tasks {
//		go func(t Task) {
//			time.Sleep(task.Duration)
//			ch <-
//		}(task)
//	}
//
//
//}
//
//
//Функция RunParallelTasks должна принимать срез задач и запускать их параллельно.
//Каждая задача должна быть запущена в отдельной горутине.
//После завершения выполнения всех задач, нужно вывести их имена и общее время, затраченное на выполнение всех задач.
//
//*Пример использования:*
//tasks := []Task{
//{ID: 1, Name: "Task 1", Duration: 2 * time.Second},
//{ID: 2, Name: "Task 2", Duration: 3 * time.Second},
//{ID: 3, Name: "Task 3", Duration: 4 * time.Second},
//}
//
//RunParallelTasks(tasks)
