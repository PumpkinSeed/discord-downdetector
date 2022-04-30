package runner

import (
	"log"

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

// TODO
// I could remove the type from Check
func Run() {
	log.Printf("[RUNNING] Downdetector")

	for i := range env.Configuration().Checks {
		handler.Handle(env.Configuration().Checks[i])
	}


	

}
