package utils

import (
	"encoding/json"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/infiniteloopcloud/discord-downdetector/env"
)

var session *discordgo.Session
var channelsCache map[string]string

func GetEvent(raw []byte) (string, error) {
	var static env.Static
	err := json.Unmarshal(raw, &static)
	if err != nil {
		return "", err
	}
	return static.Checks[0].Type, nil
}

func GetChannelID(name string) string {
	if channelsCache == nil {
		channelsCache = make(map[string]string)
	}
	if id, ok := channelsCache[name]; ok {
		return id
	} else {
		channels, err := GetSession().GuildChannels(env.Configuration().BotGuild)
		if err != nil {
			log.Print(err)
		}
		for _, channel := range channels {
			if name == channel.Name {
				channelsCache[channel.Name] = channel.ID
				return channel.ID
			}
		}
	}
	return ""
}

func GetSession() *discordgo.Session {
	if session == nil {
		var err error
		session, err = discordgo.New("Bot " + env.Configuration().BotToken)
		if err != nil {
			log.Printf("[ERROR] %s", err.Error())
		}
	}
	return session
}
