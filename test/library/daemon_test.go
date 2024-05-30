package library

import (
	"fmt"
	"testing"
	"time"

	"github.com/robfig/cron/v3"
)

func TestXxx(t *testing.T) {
	c := cron.New()
	c.AddFunc("@every 1s", func() {
		fmt.Println("tick every 1 second")
	})
	c.AddFunc("@hourly", func() {
		fmt.Println("Every hour")
	})

	c.AddFunc("@daily", func() {
		fmt.Println("Every day on midnight")
	})

	c.AddFunc("@weekly", func() {
		fmt.Println("Every week")
	})

	c.Start()

	for {
		time.Sleep(time.Second)
	}
}
