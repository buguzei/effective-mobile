package errors

const (
	ErrUserNotFound = "user not found"
)

type RepoError struct {
	Message string `json:"message"`
}

func (e *RepoError) Error() string {
	return e.Message
}
