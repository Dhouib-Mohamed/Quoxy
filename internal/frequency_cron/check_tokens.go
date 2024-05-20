package frequency_cron

import (
	database2 "api-authenticator-proxy/internal/database"
	"api-authenticator-proxy/internal/models"
	"api-authenticator-proxy/util/log"
	"strconv"
	"strings"
	"time"
)

var tokenDB = database2.Token{}
var subscriptionDB = database2.Subscription{}

func checkTokens() {
	log.Info("Checking Tokens")
	tokens, _ := tokenDB.GetAll()
	log.Info("Checking Tokens Starting")
	//log.Error(err)
	currentTime := time.Now()
	for _, token := range tokens {
		go processToken(token, currentTime)
	}
}

func processToken(token models.TokenModel, currentTime time.Time) {
	subscription, _ := subscriptionDB.GetByName(token.Subscription)
	//log.Error(err)
	frequency := subscription.Frequency
	if cronEqDate(frequency, currentTime) {
		_ = tokenDB.ResetUsage(token.Id)
		//log.Error(err)
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