package errors

type AppError struct {
	Code       string
	Message    string
	HTTPStatus int
	Err        error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return e.Message
}

func New(code, message string, status int, err error) *AppError {
	return &AppError{
		Code:       code,
		Message:    message,
		HTTPStatus: status,
		Err:        err,
	}
}
