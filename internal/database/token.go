package database

import (
	"api-authenticator-proxy/internal/models"
	"api-authenticator-proxy/internal/util/token_handler"
	"api-authenticator-proxy/util/error_handler"
	tokenError "api-authenticator-proxy/util/error_handler/token"
	"api-authenticator-proxy/util/log"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Token struct{}

var s = Subscription{}

func (t *Token) Create(token *models.CreateToken) (models.ReturnToken, error_handler.StatusError) {
	subscription, err := s.GetByName(token.Subscription)
	if err != nil {
		return models.ReturnToken{}, err
	}
	res, err1 := db.Exec("INSERT INTO token (subscription_id,passphrase) VALUES (?, ?)", subscription.Id, token.Passphrase)
	err = checkWriteResponse(res, err1, "token")
	if err != nil {
		return models.ReturnToken{}, err
	}
	id, err := GetLastInsertedId("token")
	if err != nil {
		return models.ReturnToken{}, err
	}
	resToken, err := t.GenerateToken(id, token.Passphrase)
	if err != nil {
		return models.ReturnToken{}, err
	}
	return models.ReturnToken{Token: resToken, Id: id}, nil
}

func (t *Token) GetById(id string) (models.FullToken, error_handler.StatusError) {
	var token models.FullToken
	row := db.QueryRow("SELECT token.id, subscription.name, token.current_usage, subscription.rate_limit, subscription.frequency, token.passphrase FROM token JOIN subscription ON token.subscription_id = subscription.id WHERE token.id = ?", id)
	err := checkReadResponse(row.Scan(&token.Id, &token.Subscription, &token.CurrentUsage, &token.MaxUsage, &token.Frequency, &token.Passphrase), "token")
	tokenVal, err := token_handler.Generate(token.Id, token.Passphrase)
	if err != nil {
		return models.FullToken{}, err
	}
	token.Token = tokenVal
	return token, err
}

func (t *Token) GetAll() ([]models.TokenModel, error_handler.StatusError) {
	var tokens []models.TokenModel
	rows, err := db.Query("SELECT token.id, token.subscription_id, token.current_usage, token.passphrase FROM token ORDER BY token.id ASC")
	if err != nil {
		return []models.TokenModel{}, checkReadResponse(err, "token")
	}
	for rows.Next() {
		var token models.TokenModel
		err := rows.Scan(&token.Id, &token.Subscription, &token.CurrentUsage, &token.Passphrase)
		if err != nil {
			return []models.TokenModel{}, checkReadResponse(err, "token")
		}
		tokens = append(tokens, token)
	}
	log.Debug(fmt.Sprintf("Successfully read from the token table"))
	return tokens, nil
}

func (t *Token) GetAllFull() ([]models.FullToken, error_handler.StatusError) {
	var tokens []models.FullToken
	rows, err := db.Query("SELECT token.id, subscription.name, token.current_usage, token.passphrase, subscription.rate_limit, subscription.frequency FROM token JOIN subscription ON token.subscription_id = subscription.id ORDER BY token.id ASC")
	if err != nil {
		return []models.FullToken{}, checkReadResponse(err, "token")
	}
	for rows.Next() {
		var token models.FullToken
		err := rows.Scan(&token.Id, &token.Subscription, &token.CurrentUsage, &token.Passphrase, &token.MaxUsage, &token.Frequency)
		if err != nil {
			return []models.FullToken{}, checkReadResponse(err, "token")
		}
		tokens = append(tokens, token)
	}
	log.Debug(fmt.Sprintf("Successfully read from the token table"))
	return tokens, nil
}

func (t *Token) Update(id string, token *models.UpdateToken) error_handler.StatusError {
	subscription, err := s.GetByName(token.Subscription)
	if err != nil {
		return err
	}
	res, err1 := db.Exec("UPDATE token SET subscription_id = ? WHERE id = ?", subscription.Id, id)
	return checkWriteResponse(res, err1, "token")
}

func (t *Token) Use(token string) error_handler.StatusError {
	id, err := token_handler.Decrypt(token)
	if err != nil {
		return err
	}
	tokenData, err := t.GetById(id)
	if err != nil {
		return err
	}
	if tokenData.MaxUsage <= tokenData.CurrentUsage {
		return tokenError.LimitedTokenError()
	}
	res, err1 := db.Exec("UPDATE token SET current_usage = current_usage + 1 WHERE id = ?", id)
	return checkWriteResponse(res, err1, "token")
}

func (t *Token) Disable(id string) error_handler.StatusError {
	res, err := db.Exec("DELETE FROM token WHERE id = ?", id)
	return checkWriteResponse(res, err, "token")
}

func (t *Token) GenerateToken(id string, passphrase string) (string, error_handler.StatusError) {
	return token_handler.Generate(id, passphrase)
}

func (t *Token) ResetUsage(ids []string) error_handler.StatusError {
	query, args, err := sqlx.In(`UPDATE token SET current_usage = 0 WHERE id IN (?) AND current_usage != 0`, ids)
	if err != nil {
		return error_handler.UnexpectedError(fmt.Sprintf("Error when creating query: %v", err))
	}

	res, err := db.Exec(query, args...)
	return checkWriteResponse(res, err, "token")
}
