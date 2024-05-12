package usecase

import (
	"context"
	"github.com/buguzei/effective-mobile/auth/internal/models"
	"github.com/buguzei/effective-mobile/auth/internal/repo/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

// positive test for AuthUC.SignUp
func TestAuthUC_SignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mock_repo.NewMockUserRepo(ctrl)
	refreshRepo := mock_repo.NewMockRefreshRepo(ctrl)

	user := models.User{
		ID:       99,
		Email:    "test@example.com",
		Name:     "Richard Hendricks",
		Password: "password",
	}
	ctx := context.Background()

	userRepo.EXPECT().IsUserExist(ctx, user.Email).Return(false, nil).Times(1)

	userRepo.EXPECT().NewUser(ctx, user).Return(user.ID, nil).Times(1)

	refreshRepo.EXPECT().SetRefresh(ctx, gomock.Any(), user.ID).Return(nil).Times(1)

	uc := New(refreshRepo, userRepo)
	_, err := uc.SignUp(ctx, user)

	require.NoError(t, err)
}

// negative test for AuthUC.SignUp
//func TestAuthUC_SignUpErr(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	userRepo := mock_repo.NewMockUserRepo(ctrl)
//	refreshRepo := mock_repo.NewMockRefreshRepo(ctrl)
//
//}

// positive test for AuthUC.SignIn
func TestAuthUC_SignIn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mock_repo.NewMockUserRepo(ctrl)
	refreshRepo := mock_repo.NewMockRefreshRepo(ctrl)

	user := models.User{
		ID:       123,
		Name:     "Gavin Belson",
		Email:    "good@example.com",
		Password: "sui",
	}

	ctx := context.Background()

	userRepo.EXPECT().VerifyEmailAndPass(ctx, user.Email, user.Password).Return(user, nil).Times(1)

	refreshRepo.EXPECT().SetRefresh(ctx, gomock.Any(), user.ID).Return(nil).Times(1)

	uc := New(refreshRepo, userRepo)
	_, err := uc.SignIn(ctx, user)

	require.NoError(t, err)
}
