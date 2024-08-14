package tests

import (
	"api-authenticator-proxy/internal/database"
	"api-authenticator-proxy/internal/models"
	"testing"
	"time"
)

type TestToken struct {
	t  *testing.T
	tk *database.Token
}

func (tt *TestToken) create(token models.CreateToken, status int) models.ReturnToken {
	tk, res := tt.tk.Create(&token)
	validateError(tt.t, res, status)
	if status == 200 {
		return tk
	}
	return models.ReturnToken{}
}

func (tt *TestToken) getById(id string, status int) {
	_, err := tt.tk.GetById(id)
	validateError(tt.t, err, status)
}

func (tt *TestToken) getAll(status int) {
	_, err := tt.tk.GetAll()
	validateError(tt.t, err, status)
}

func (tt *TestToken) getAllFull(status int) {
	_, err := tt.tk.GetAllFull()
	validateError(tt.t, err, status)
}

func (tt *TestToken) update(id string, token models.UpdateToken, status int) {
	err := tt.tk.Update(id, &token)
	validateError(tt.t, err, status)
}

func (tt *TestToken) disable(id string, status int) {
	err := tt.tk.Disable(id)
	validateError(tt.t, err, status)
}

func (tt *TestToken) use(tk string, status int) {
	err := tt.tk.Use(tk)
	validateError(tt.t, err, status)
}

func TestTokenWorkflow(t *testing.T) {
	t.Run("Initialize DB", testDatabase)

	var token models.ReturnToken

	var usedToken models.FullToken

	t.Run("Create", func(t *testing.T) {
		ts := TestSubscription{t: t, s: &database.Subscription{}}
		name := time.Now().String()
		testCorrectSubscription := models.CreateSubscription{
			Name:      name,
			Frequency: "* * * * *",
			RateLimit: 2,
		}
		testToken := models.CreateToken{
			Subscription: name,
			Passphrase:   "test",
		}
		testInvalidToken := models.CreateToken{
			Subscription: "test1",
			Passphrase:   "test",
		}
		ts.create(testCorrectSubscription, 200)
		tt := TestToken{t: t, tk: &database.Token{}}
		token = tt.create(testToken, 200)
		tt.create(testInvalidToken, 404)
	})
	t.Run("GetById", func(t *testing.T) {
		tt := TestToken{t: t, tk: &database.Token{}}
		tt.getById(token.Id, 200)
		tt.getById("2", 404)
	})
	t.Run("Update", func(t *testing.T) {
		tt := TestToken{t: t, tk: &database.Token{}}
		tt.getById(token.Id, 200)
		tt.getById("2", 404)
		tt.update(token.Id, models.UpdateToken{Subscription: "test2"}, 200)
	})
	t.Run("Use", func(t *testing.T) {
		tt := TestToken{t: t, tk: &database.Token{}}
		tt.use(token.Token, 200)
		tt.use("2", 401)
		usedToken, _ = tt.tk.GetById(token.Id)
		if usedToken.CurrentUsage != 1 {
			t.Errorf("Expected 1, got %d", usedToken.CurrentUsage)
		}
	})
	t.Run("Disable", func(t *testing.T) {
		tt := TestToken{t: t, tk: &database.Token{}}
		tt.disable(token.Id, 200)
		tt.getById(token.Id, 404)
		tt.disable("2", 404)
	})
}

func TestTokenGetAll(t *testing.T) {
	t.Run("Initialize DB", testDatabase)
	tt := TestToken{t: t, tk: &database.Token{}}
	tt.getAll(200)
}

func TestTokenGetAllFull(t *testing.T) {
	t.Run("Initialize DB", testDatabase)
	tt := TestToken{t: t, tk: &database.Token{}}
	tt.getAllFull(200)
}
