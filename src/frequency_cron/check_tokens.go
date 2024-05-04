package frequency_cron

import (
	"api-authenticator-proxy/src/database"
	"api-authenticator-proxy/src/database/models"
	"strconv"
	"strings"
	"time"
)

var tokenDB = database.Token{}
var subscriptionDB = database.Subscription{}

func check_tokens() {
	tokens, err := tokenDB.GetAll()
	if err != nil {
		panic(err)
	}
	currentTime := time.Now()
	for _, token := range tokens {
		go processToken(token, currentTime)
	}
}

func processToken(token models.TokenModel, currentTime time.Time) {
	subscription, err := subscriptionDB.GetByName(token.Subscription)
	if err != nil {
		panic(err)
	}
	frequency := subscription.Frequency
	if cronEqDate(frequency, currentTime) {
		err = tokenDB.ResetUsage(token.Id)
		if err != nil {
			panic(err)
		}
	}
}

func cronEqDate(frequency string, currentTime time.Time) bool {
	crontime := strings.Split(frequency, " ")
	datetime := []int{currentTime.Minute(), currentTime.Hour(), currentTime.Day(), int(currentTime.Month()), currentTime.Year()}
	for i, v := range crontime {
		if v != "*" {
			n, e := strconv.Atoi(v)
			if e != nil {
				panic(e)
			}
			if n != datetime[i] {
				return false
			}
		}
	}
	return true
}
