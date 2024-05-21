package frequency_cron

import (
	database2 "api-authenticator-proxy/internal/database"
	"api-authenticator-proxy/internal/models"
	"api-authenticator-proxy/util/log"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

var tokenDB = database2.Token{}
var subscriptionDB = database2.Subscription{}

func checkTokens() {
	log.Debug("Checking Tokens...")

	tokens, err := tokenDB.GetAllFull()
	var wg sync.WaitGroup
	validTokensChan := make(chan string, len(tokens))

	if err != nil {
		_, err := err.GetError()
		log.Error(fmt.Errorf("error getting tokens: %v", err))
		return
	}
	currentTime := time.Now()
	for _, token := range tokens {
		wg.Add(1)
		go processToken(token, currentTime, validTokensChan, &wg)
	}
	go func() {
		wg.Wait()
		close(validTokensChan)
	}()

	// Collect valid tokens from the channel
	var validTokens []string
	for token := range validTokensChan {
		validTokens = append(validTokens, token)
	}
	tokenDB.ResetUsage(validTokens)
}

func processToken(token models.FullToken, currentTime time.Time, validTokensChan chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	frequency := token.Frequency
	if cronEqDate(frequency, currentTime) {
		validTokensChan <- token.Id
	}
}

func cronEqDate(frequency string, currentTime time.Time) bool {
	crontime := strings.Split(frequency, " ")
	datetime := []int{currentTime.Minute(), currentTime.Hour(), currentTime.Day(), int(currentTime.Month()), currentTime.Year()}
	for i, v := range crontime {
		if v != "*" {
			n, e := strconv.Atoi(v)
			log.Fatal(e)
			if n != datetime[i] {
				return false
			}
		}
	}
	return true
}
