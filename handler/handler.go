package handler

import (
	"log"
	"net"
	"time"

	embed "github.com/Clinet/discordgo-embed"
	"github.com/bwmarrin/discordgo"
	env "github.com/infiniteloopcloud/discord-downdetector/env"
)

const (
	warning = 0xD10000
)

var channelName string

func Handle(body env.Check) (string, *discordgo.MessageEmbed, error) {
	if !checkHealth(body) {
		return unreachable(body)
	} else {
		return "", nil, nil
	}
}

func unreachable(check env.Check) (string, *discordgo.MessageEmbed, error) {

	message := embed.NewEmbed().
		SetAuthor("Port " + check.Port).
		SetTitle("[Host unreachable] " + check.Value).
		SetColor(warning)

	return env.Configuration().ChannelName, message.MessageEmbed, nil
}

func checkHealth(check env.Check) bool {
	timeout := 1 * time.Second

	conn, err := net.DialTimeout("tcp", net.JoinHostPort(check.Value, check.Port), timeout)
	if err != nil {
		log.Println("[ERROR] unreachable", check.Value+":"+check.Port)
	}
	if conn != nil {
		defer conn.Close()
		return true
	}

	return false
}
