package models

type SubscriptionModel struct {
	Id         string
	Name       string
	Frequency  string
	RateLimit  int
	Deprecated bool
}

type CreateSubscription struct {
	Name      string `name:"name" binding:"required"`
	Frequency string `name:"frequency" binding:"required"`
	RateLimit int    `name:"rateLimit" binding:"required"`
}

type UpdateSubscription struct {
	Name      string `name:"name"`
	Frequency string `name:"frequency"`
	RateLimit int    `name:"rateLimit"`
}
