package token

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"os"
)

const (
	refreshTokenSize = 32
)

type Pair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewAccessToken(userID int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS512)

	claims := token.Claims.(jwt.MapClaims)

	claims["sub"] = userID

	strToken, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", fmt.Errorf("error of signing: %w", err)
	}

	return strToken, nil
}

func NewRefreshToken() (string, error) {
	rb := make([]byte, refreshTokenSize)

	_, err := rand.Read(rb)
	if err != nil {
		return "", fmt.Errorf("error of reading: %w", err)
	}

	return base64.URLEncoding.EncodeToString(rb), nil
}
