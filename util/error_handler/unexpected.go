package error_handler

type Unexpected struct {
	Message string
	Code    int
}

func (e Unexpected) GetError() (int, string) {
	return e.Code, "Unexpected error: " + e.Message
}

func UnexpectedError(message string) Unexpected {
	return Unexpected{Message: message, Code: 500}
}
