package runner

import (
	"log"
	"time"

	"github.com/infiniteloopcloud/discord-downdetector/env"
	handler "github.com/infiniteloopcloud/discord-downdetector/handler"
	utils "github.com/infiniteloopcloud/discord-downdetector/utils"
)

func check(body env.Check) {
	channel, message, err := handler.Handle(body)
	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return
	}
	channelID := utils.GetChannelID(channel)
	if channelID == "" {
		channelID = utils.GetChannelID("unknown")
	}
	if channelID != "" && message != nil {
		_, err = utils.GetSession().ChannelMessageSendEmbed(channelID, message)
		if err != nil {
			log.Printf("[ERROR] %s", err.Error())
		}
	}
}

func Run() {
	log.Printf("[RUNNING] Downdetector")

	// A loop what runs forever to check is the host reachable
	for {
		for i := range env.Configuration().Checks {
			check(env.Configuration().Checks[i])
		}
		// Checks only the first object's interval
		// Don't wait between objects
		interval, unit := utils.GetTime(env.Configuration().Checks[0].Interval)
		switch unit {
		case "h":
			time.Sleep(time.Duration(interval) * time.Hour)
		case "m":
			time.Sleep(time.Duration(interval) * time.Minute)
		case "s":
			time.Sleep(time.Duration(interval) * time.Second)
		}
	}

}
