package utils

type ErrorRes struct {
	Error string `json:"error"`
}

func NewError(err error) ErrorRes {
	errRes := ErrorRes{Error: err.Error()}

	return errRes
}
