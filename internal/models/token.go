package models

type TokenModel struct {
	Id           string
	Passphrase   string
	Subscription string
	CurrentUsage int
}

type FullToken struct {
	Id           string
	Passphrase   string
	Subscription string
	CurrentUsage int
	MaxUsage     int
	Frequency    string
	Token        string
}

type CreateToken struct {
	Passphrase   string `name:"passphrase"`
	Subscription string `name:"subscription" binding:"required"`
}

type UpdateToken struct {
	Subscription string `name:"subscription" binding:"required"`
}

type ReturnToken struct {
	Token string `json:"token"`
	Id    string `json:"id"`
}
