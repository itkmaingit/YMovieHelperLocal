package utils

type CustomError struct {
	FrontMsg string
	BackMsg  string
}

func (e *CustomError) Error() string {
	return e.BackMsg
}

func (e *CustomError) FrontError() string {
	return e.FrontMsg
}
