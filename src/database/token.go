package database

import (
	"database/sql"
	"strconv"
)

type TokenModel struct {
	id           string
	passphrase   string
	subscription Subscription
	currentUsage int
}

type CreateToken struct {
	passphrase   string
	subscription string
}

type UpdateToken struct {
	subscription string
}

type Token struct{}

func (t *Token) Create(token *CreateToken) (string, error) {
	s := Subscription{}
	subscription, err := s.GetByName(token.subscription)
	if err != nil {
		return "", err
	}
	res, err := db.Exec("INSERT INTO token (passphrase, subscription) VALUES (?, ?)", token.passphrase, subscription.id)
	if err != nil {
		return "", err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(id, 10), nil
}

func (t *Token) GetById(id string) (TokenModel, error) {
	var token TokenModel
	row := db.QueryRow("SELECT token.id, subscription.name, token.currentUsage FROM token JOIN subscription ON token.subscription = subscription.id WHERE token.id = ?", id)
	err := row.Scan(&token.id, &token.subscription, &token.currentUsage)
	return token, err
}

func (t *Token) GetAll() ([]TokenModel, error) {
	var tokens []TokenModel
	rows, err := db.Query("SELECT * FROM token")
	if err != nil {
		return []TokenModel{}, err
	}
	for rows.Next() {
		var token TokenModel
		err := rows.Scan(&token.id, &token.subscription, &token.currentUsage)
		if err != nil {
			return []TokenModel{}, err
		}
		tokens = append(tokens, token)
	}
	return tokens, nil
}

func (t *Token) Update(id string, token *UpdateToken) (sql.Result, error) {
	s := Subscription{}
	subscription, err := s.GetByName(token.subscription)
	if err != nil {
		return nil, err
	}
	return db.Exec("UPDATE token SET subscription = ? WHERE id = ?", id, subscription.id, id)
}

// TODO : check if token reached its limit yet
func (t *Token) Use(token string) (sql.Result, error) {

	//token, err := t.GetById(id)
	//if err!=nil {
	//	return nil, err
	//}
	//subscriptionId, err := token_handler.Decrypt(token.value)
	//if err!=nil {
	//	return nil, err
	//}
	//subscription :=
	//return db.Exec("UPDATE token SET currentUsage = currentUsage + 1 WHERE id = ?", id)
	return db.Exec("UPDATE token SET currentUsage = currentUsage + 1 WHERE id = ?", token)
}

func (t *Token) Disable(id string) (sql.Result, error) {
	return db.Exec("DELETE FROM token WHERE id = ?", id)
}
