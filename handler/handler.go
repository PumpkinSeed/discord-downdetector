package handler

import (
	"log"
	"net/http"
	"strconv"

	embed "github.com/Clinet/discordgo-embed"
	"github.com/bwmarrin/discordgo"
	env "github.com/infiniteloopcloud/discord-downdetector/env"
)

const (
	warning = 0xD10000
)

var channelName string

// Check is the endpoint alive
func Handle(body env.Check) (string, *discordgo.MessageEmbed, error) {
	var statusCode = http.StatusOK
	code := checkHealth(body)

	if body.Parameters != nil && body.Parameters.StatusCode != 0 {
		statusCode = body.Parameters.StatusCode
	}
	if code != statusCode {
		return unreachable(body, code)
	} else {
		return "", nil, nil
	}

}

// Send an embed to the downdetector channel
func unreachable(check env.Check, code int) (string, *discordgo.MessageEmbed, error) {
	status := strconv.Itoa(code)
	message := embed.NewEmbed().
		SetAuthor("Status code: " + status).
		SetTitle("[Host unreachable] " + check.Value).
		SetColor(warning)

	return env.Configuration().ChannelName, message.MessageEmbed, nil
}

// Return the status code of the request
func checkHealth(check env.Check) int {
	resp, err := http.Get(check.Type + "://" + check.Value)
	if err != nil {
		log.Println("[ERROR]", err)
		return resp.StatusCode
	}
	defer resp.Body.Close()

	return resp.StatusCode
}
