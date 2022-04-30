package handler

import (
	"log"
	"net"
	"strconv"
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
		SetAuthor(check.Type).
		SetTitle("["+check.Value+"] Host unreachable").
		SetColor(warning)

	return check.ChannelName, message.MessageEmbed, nil
}

func checkHealth(check env.Check) bool {
	timeout := 1 * time.Second
	port := strconv.Itoa(check.Port)

	conn, err := net.DialTimeout("tcp", net.JoinHostPort(check.Value, port), timeout)
	if err != nil {
		log.Println("Connecting error:", err)
	}
	if conn != nil {
		defer conn.Close()
		log.Println("[DEBUG] Reachable", net.JoinHostPort(check.Value, port))
		return true
	}

	return false
}
