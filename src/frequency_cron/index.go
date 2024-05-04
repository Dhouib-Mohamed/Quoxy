package frequency_cron

import "github.com/robfig/cron"

func Init() {
	c := cron.New()
	err := c.AddFunc("0 * * * *", check_tokens)
	if err != nil {
		panic(err)
	}
	c.Start()
}
