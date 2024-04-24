package services

import (
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

func StartCronJob() {
	c := cron.New(
		cron.WithLocation(time.UTC),
	)

	if _, err := c.AddFunc("0 * * * *", func() { //every 1 hour
		BuyBitcoin()
	}); err != nil {
		log.Fatalln(err.Error())
	}

	c.Start()
	fmt.Println("cron job started...")
}
