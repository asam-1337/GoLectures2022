package errors

import "fmt"

type HttpError struct {
	Code    int
	Message string
}

func (e HttpError) Error() string {
	return fmt.Sprintf(`"status": "%s", "msg": "%s"`, e.Code, e.Message)
}

func HttpErrorResponse(code int, msg string) string {
	return fmt.Sprintf(`"status": "%s", "msg": "%s"`, code, msg)
}
