package pkg

import (
	"ecommerce-backend/src/configs"
	"encoding/json"

	"github.com/bwmarrin/discordgo"
)

type ServerEnvironment struct {
	Hostname string
	Url      string
	Method   string
}

type DiscordErrorLog struct {
	Level     string `json:"level"`
	Caller    string `json:"caller"`
	Message   string `json:"msg"`
	Timestamp string `json:"timestamp"`
}

func WebhookSend(text string, ev ServerEnvironment) error {
	data := DiscordErrorLog{}

	if err := json.Unmarshal([]byte(text), &data); err != nil {
		// log.Printf("Discord webhook error: %v", err)
		return err
	}

	emb := []*discordgo.MessageEmbed{
		{
			Type:        "rich",
			Title:       "Environment",
			Description: "**End Point**\n" + ev.Hostname,
			Color:       15548997,
			Timestamp:   data.Timestamp,
		},
	}

	s := "\n**[" + ev.Method + "]** `" + ev.Url + "`"

	hookMessage := &discordgo.WebhookParams{
		Embeds:  emb,
		Content: "**Server Error :boom:** - TypeError: " + data.Message + " at " + data.Caller + s,
	}

	dc, _ := discordgo.New("Bot")

	_, err := dc.WebhookExecute(configs.Cfg.DiscordWebhook.ID, configs.Cfg.DiscordWebhook.Token, false, hookMessage)
	if err != nil {
		// log.Printf("Discord webhook error: %v", err)
		return err
	}
	return nil
}
