package http

import "net/http"

func (h Handler) VerifyToken(next http.Handler) http.Handler {
	return nil
}
