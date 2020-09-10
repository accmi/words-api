package utils

type ErrorsResponse struct {
	Error string `json:"error"`
}

func (er *ErrorsResponse) GetError(es string) {
	er.Error = es
}

type ErrorsResponseInterface interface {
	GetError()
}
