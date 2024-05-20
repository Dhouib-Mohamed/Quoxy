package token

type LimitedToken struct {
	Code int
}

func (e LimitedToken) GetError() (int, string) {
	return e.Code, "This token has reached its limit, please try again later"
}

func LimitedTokenError() LimitedToken {
	return LimitedToken{Code: 401}
}
