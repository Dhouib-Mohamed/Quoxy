package database

import (
	"api-authenticator-proxy/src/database/models"
	"fmt"
	"regexp"
	"strings"
	"time"
)

type Subscription struct{}

func (s *Subscription) Create(subscription *models.CreateSubscription) error {
	frequency, err := validateFrequency(subscription.Frequency)
	if err != nil {
		return err
	}
	return checkResponse(db.Exec("INSERT INTO subscription (name, frequency, rate_limit) VALUES (?, ?, ?)", subscription.Name, frequency, subscription.RateLimit))
}

func (s *Subscription) GetByName(name string) (models.SubscriptionModel, error) {
	var subscription models.SubscriptionModel
	row := db.QueryRow("SELECT id, name, frequency, rate_limit, deprecated FROM subscription WHERE name = ?", name)
	err := row.Scan(&subscription.Id, &subscription.Name, &subscription.Frequency, &subscription.RateLimit, &subscription.Deprecated)
	if subscription.Deprecated {
		return models.SubscriptionModel{}, fmt.Errorf("subscription is out of service")
	}
	return subscription, err
}

func (s *Subscription) GetById(id string) (models.SubscriptionModel, error) {
	var subscription models.SubscriptionModel
	row := db.QueryRow("SELECT id, name, frequency, rate_limit, deprecated  FROM subscription WHERE id = ?", id)
	err := row.Scan(&subscription.Id, &subscription.Name, &subscription.Frequency, &subscription.RateLimit, &subscription.Deprecated)
	return subscription, err
}

func (s *Subscription) Update(id string, subscription *models.UpdateSubscription) error {
	return checkResponse(db.Exec("UPDATE subscription SET name = ?, frequency = ?, rate_limit = ? WHERE id = ?", subscription.Name, subscription.Frequency, subscription.RateLimit, id))
}

func (s *Subscription) Disable(id string) error {
	return checkResponse(db.Exec("UPDATE subscription SET deprecated = true WHERE id = ?", id))
}

func (s *Subscription) Restore(id string) error {
	return checkResponse(db.Exec("UPDATE subscription SET deprecated = false WHERE id = ?", id))
}

func (s *Subscription) GetAll() ([]models.SubscriptionModel, error) {
	var subscriptions []models.SubscriptionModel
	rows, err := db.Query("SELECT id, name, frequency, rate_limit FROM subscription where deprecated = false")
	if err != nil {
		return []models.SubscriptionModel{}, err
	}
	for rows.Next() {
		var subscription models.SubscriptionModel
		err = rows.Scan(&subscription.Id, &subscription.Name, &subscription.Frequency, &subscription.RateLimit)
		if err != nil {
			return []models.SubscriptionModel{}, err
		}
		subscriptions = append(subscriptions, subscription)
	}
	return subscriptions, nil
}

func validateFrequency(frequency string) (string, error) {
	frequency = strings.Trim(frequency, " ")
	currentTime := time.Now()
	lowerFreq := strings.ToLower(frequency)

	if lowerFreq == "every-minute" || lowerFreq == "everyminute" || lowerFreq == "every minute" || lowerFreq == "every_minute" {
		return "* * * * *", nil
	}
	if lowerFreq == "hourly" {
		return fmt.Sprintf("%d * * * *", currentTime.Minute()), nil
	}
	if lowerFreq == "daily" {
		return fmt.Sprintf("%d %d * * *", currentTime.Minute(), currentTime.Hour()), nil
	}
	if lowerFreq == "monthly" {
		return fmt.Sprintf("%d %d %d * *", currentTime.Minute(), currentTime.Hour(), currentTime.Day()), nil
	}
	if lowerFreq == "yearly" {
		return fmt.Sprintf("%d %d %d %d *", currentTime.Minute(), currentTime.Hour(), currentTime.Day(), currentTime.Month()), nil
	}

	regexString := `^(\*|[0-5]?[0-9]) (\*|[0-5]?[0-9]) (\*|[01]?[0-9]|2[0-3]) (\*|[1-9]|[12][0-9]|3[01]) (\*|[1-9]|1[0-2])$`
	match, err := regexp.MatchString(regexString, frequency)
	if err != nil {
		return "", err
	}
	if match {
		return frequency, nil
	}
	return "", fmt.Errorf("invalid frequency, must be daily, monthly or a cron expression of 3 numbers or * separated by spaces")
}
