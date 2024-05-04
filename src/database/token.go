package database

import (
	"api-authenticator-proxy/src/database/models"
	"api-authenticator-proxy/src/utils/token_handler"
	"fmt"
)

type Token struct{}

var s = Subscription{}

func (t *Token) Create(token *models.CreateToken) (models.ReturnToken, error) {
	subscription, err := s.GetByName(token.Subscription)
	if err != nil {
		return models.ReturnToken{}, err
	}
	err = checkResponse(db.Exec("INSERT INTO token (subscription_id,passphrase) VALUES (?, ?)", subscription.Id, token.Passphrase))
	if err != nil {
		fmt.Println("Error", err)
		return models.ReturnToken{}, err
	}
	id, err := GetLastInsertedId("token")
	if err != nil {
		return models.ReturnToken{}, err
	}
	res, err := t.GenerateToken(id, token.Passphrase)
	if err != nil {
		return models.ReturnToken{}, err
	}
	return models.ReturnToken{Token: res, Id: id}, nil
}

func (t *Token) GetById(id string) (models.TokenModel, error) {
	var token models.TokenModel
	row := db.QueryRow("SELECT token.id, subscription.name, token.current_usage FROM token JOIN subscription ON token.subscription_id = subscription.id WHERE token.id = ?", id)
	err := row.Scan(&token.Id, &token.Subscription, &token.CurrentUsage)
	return token, err
}

func (t *Token) GetAll() ([]models.TokenModel, error) {
	var tokens []models.TokenModel
	rows, err := db.Query("SELECT token.id, subscription.name, token.current_usage FROM token JOIN subscription ON token.subscription_id = subscription.id ORDER BY token.id ASC")
	if err != nil {
		return []models.TokenModel{}, err
	}
	for rows.Next() {
		var token models.TokenModel
		err := rows.Scan(&token.Id, &token.Subscription, &token.CurrentUsage)
		if err != nil {
			return []models.TokenModel{}, err
		}
		tokens = append(tokens, token)
	}
	return tokens, nil
}

func (t *Token) Update(id string, token *models.UpdateToken) error {
	subscription, err := s.GetByName(token.Subscription)
	if err != nil {
		return err
	}
	return checkResponse(db.Exec("UPDATE token SET subscription_id = ? WHERE id = ?", subscription.Id, id))
}

func (t *Token) Use(token string) error {
	id, err := token_handler.Decrypt(token)
	if err != nil {
		return err
	}
	tokenData, err := t.GetById(id)
	if err != nil {
		return err
	}
	fmt.Println(tokenData.Subscription)
	subscription, err := s.GetByName(tokenData.Subscription)
	if err != nil {
		return err
	}
	if subscription.RateLimit <= tokenData.CurrentUsage {
		return fmt.Errorf("token has reached its limit")
	}
	return checkResponse(db.Exec("UPDATE token SET current_usage = current_usage + 1 WHERE id = ?", id))
}

func (t *Token) Disable(id string) error {
	return checkResponse(db.Exec("DELETE FROM token WHERE id = ?", id))
}

func (t *Token) GenerateToken(id string, passphrase string) (string, error) {
	return token_handler.Generate(id, passphrase)
}

func (t *Token) ResetUsage(id string) error {
	return checkResponse(db.Exec("UPDATE token SET current_usage = 0 WHERE id = ?", id))
}
