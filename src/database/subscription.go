package database

import (
	"database/sql"
	"fmt"
)

type SubscriptionModel struct {
	id         string
	name       string
	frequency  string
	rateLimit  int
	deprecated bool
}

type CreateSubscription struct {
	name      string
	frequency string
	rateLimit int
}

type UpdateSubscription struct {
	name      string
	frequency string
	rateLimit int
}

type Subscription struct{}

func (s *Subscription) Create(subscription *CreateSubscription) (sql.Result, error) {
	return db.Exec("INSERT INTO subscription (name, frequency, rateLimit) VALUES (?, ?, ?)", subscription.name, subscription.frequency, subscription.rateLimit)
}

func (s *Subscription) GetByName(name string) (SubscriptionModel, error) {
	var subscription SubscriptionModel
	row := db.QueryRow("SELECT * FROM subscription WHERE name = ?", name)
	err := row.Scan(&subscription.id, &subscription.name, &subscription.frequency, &subscription.rateLimit, &subscription.deprecated)
	if subscription.deprecated {
		return SubscriptionModel{}, fmt.Errorf("subscription is out of service")
	}
	return subscription, err
}

func (s *Subscription) GetById(id string) (SubscriptionModel, error) {
	var subscription SubscriptionModel
	row := db.QueryRow("SELECT * FROM subscription WHERE id = ?", id)
	err := row.Scan(&subscription.id, &subscription.name, &subscription.frequency, &subscription.rateLimit)
	return subscription, err
}

func (s *Subscription) Update(id string, subscription *UpdateSubscription) (sql.Result, error) {
	return db.Exec("UPDATE subscription SET name = ?, frequency = ?, rateLimit = ? WHERE id = ?", subscription.name, subscription.frequency, subscription.rateLimit, id)
}

func (s *Subscription) Disable(id string) (sql.Result, error) {
	return db.Exec("UPDATE subscription SET deprecated = true WHERE id = ?", id)
}

func (s *Subscription) Restore(id string) (sql.Result, error) {
	return db.Exec("UPDATE subscription SET deprecated = false WHERE id = ?", id)
}

func (s *Subscription) GetAll() ([]SubscriptionModel, error) {
	var subscriptions []SubscriptionModel
	rows, err := db.Query("SELECT * FROM subscription")
	if err != nil {
		return []SubscriptionModel{}, err
	}
	for rows.Next() {
		var subscription SubscriptionModel
		err = rows.Scan(&subscription.id, &subscription.name, &subscription.frequency, &subscription.rateLimit)
		if err != nil {
			return []SubscriptionModel{}, err
		}
		subscriptions = append(subscriptions, subscription)
	}
	return subscriptions, nil
}
