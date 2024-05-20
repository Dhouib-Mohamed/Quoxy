package token

type InvalidToken struct {
	Code int
}

func (e InvalidToken) GetError() (int, string) {
	return e.Code, "Invalid token provided"
}

func InvalidTokenError() InvalidToken {
	return InvalidToken{Code: 401}
}
