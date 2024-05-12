package frequency_cron

import (
	"api-authenticator-proxy/src/utils/log"
	"github.com/robfig/cron"
)

func Init() {
	c := cron.New()
	log.Fatal(c.AddFunc("0 * * * *", checkTokens))
	log.Info("The CRON job is starting")
	c.Start()
}
