package database

import (
	"api-authenticator-proxy/internal/models"
	"api-authenticator-proxy/util/error_handler"
	dbError "api-authenticator-proxy/util/error_handler/db"
	"api-authenticator-proxy/util/id"
	"api-authenticator-proxy/util/log"
	"fmt"
	"regexp"
	"strings"
	"time"
)

type Subscription struct{}

func (s *Subscription) Create(subscription *models.CreateSubscription) (string, error_handler.StatusError) {
	frequency, err := validateFrequency(subscription.Frequency)
	if subscription.RateLimit < 1 {
		return "", dbError.FieldConstraintError("subscription", "rate_limit", "should be greater than 0")
	}
	if err != nil {
		return "", err
	}
	generatedId := id.GenerateRandomId()
	result, err1 := db.Exec("INSERT INTO subscription (id,name, frequency, rate_limit) VALUES (?,?, ?, ?)", generatedId, subscription.Name, frequency, subscription.RateLimit)
	err = checkWriteResponse(result, err1, "subscription")
	if err != nil {
		return "", err
	}
	return generatedId, nil
}

func (s *Subscription) GetByName(name string) (models.SubscriptionModel, error_handler.StatusError) {
	var subscription models.SubscriptionModel
	log.Debug("Reading from the subscription table with name : ", name)
	row := db.QueryRow("SELECT id, name, frequency, rate_limit, deprecated FROM subscription WHERE name = ?", name)

	err := checkReadResponse(row.Scan(&subscription.Id, &subscription.Name, &subscription.Frequency, &subscription.RateLimit, &subscription.Deprecated), "subscription")
	if subscription.Deprecated {
		return models.SubscriptionModel{}, dbError.CanceledElementError("subscription")
	}
	log.Debug("Successfully read 1 item from the subscription table ", subscription)
	return subscription, err
}

func (s *Subscription) GetById(id string) (models.SubscriptionModel, error_handler.StatusError) {
	var subscription models.SubscriptionModel
	row := db.QueryRow("SELECT id, name, frequency, rate_limit, deprecated  FROM subscription WHERE id = ?", id)
	err := checkReadResponse(row.Scan(&subscription.Id, &subscription.Name, &subscription.Frequency, &subscription.RateLimit, &subscription.Deprecated), "subscription")
	return subscription, err
}

func (s *Subscription) Update(id string, subscription *models.UpdateSubscription) error_handler.StatusError {
	res, err := db.Exec("UPDATE subscription SET name = ?, frequency = ?, rate_limit = ? WHERE id = ?", subscription.Name, subscription.Frequency, subscription.RateLimit, id)
	return checkWriteResponse(res, err, "subscription")
}

// This will delete all the tokens linked to this subscription
func (s *Subscription) Disable(id string) error_handler.StatusError {
	res, err := db.Exec("DELETE FROM token where subscription_id = ?", id)
	finalError := checkWriteResponse(res, err, "token")
	if finalError != nil {
		return finalError
	}
	res, err = db.Exec("UPDATE subscription SET deprecated = true WHERE id = ? and deprecated = false", id)
	return checkWriteResponse(res, err, "subscription")
}

func (s *Subscription) Restore(id string) error_handler.StatusError {
	res, err := db.Exec("UPDATE subscription SET deprecated = false WHERE id = ? and deprecated = true", id)
	return checkWriteResponse(res, err, "subscription")
}

func (s *Subscription) GetAll() ([]models.SubscriptionModel, error_handler.StatusError) {
	var subscriptions []models.SubscriptionModel
	rows, err := db.Query("SELECT id, name, frequency, rate_limit FROM subscription where deprecated = false")
	if err != nil {
		return []models.SubscriptionModel{}, checkReadResponse(err, "subscription")
	}
	for rows.Next() {
		var subscription models.SubscriptionModel
		err = rows.Scan(&subscription.Id, &subscription.Name, &subscription.Frequency, &subscription.RateLimit)
		if err != nil {
			return []models.SubscriptionModel{}, checkReadResponse(err, "subscription")
		}
		subscriptions = append(subscriptions, subscription)
	}
	log.Debug("Successfully read ", len(subscriptions), " items from the subscription table")
	return subscriptions, nil
}

func validateFrequency(frequency string) (string, error_handler.StatusError) {
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
		return "", dbError.IncorrectFrequencyError()
	}
	if match {
		return frequency, nil
	}
	return "", dbError.IncorrectFrequencyError()
}
