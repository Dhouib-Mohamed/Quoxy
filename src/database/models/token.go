package models

type TokenModel struct {
	Id           string
	Passphrase   string
	Subscription string
	CurrentUsage int
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
