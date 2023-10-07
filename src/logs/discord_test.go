package logs_test

import (
	"ecommerce-backend/src/configs"
	"ecommerce-backend/src/logs"
	"encoding/json"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"
)

func init() {
	configs.InitConfigMock()
}

func TestWebHook(t *testing.T) {
	invalidSyntaxText := "mock"
	validSyntaxText := `{"level":"error","timestamp":"2023-10-05T21:59:42.685+0700","caller":"services/user.go:34","msg":"testing"}`
	notFoundID := "1159508574125428721"
	invalidID := "asdkasklfjafasf"
	incorrectToken := "yoyoyoyoyoyyoyooyo"
	envMock := logs.ServerEnvironment{
		Hostname: "localhost:3333",
		Url:      "/api/test",
		Method:   "GET",
	}

	t.Run("Unmarshal error", func(t *testing.T) {
		err := logs.WebhookSend(invalidSyntaxText, envMock)
		assert.Error(t, err)
		assert.ErrorIs(t, err, err.(*json.SyntaxError))
	})

	// ปิดไว้เพราะมัน send รัวๆๆๆๆๆๆๆๆๆๆๆๆ ตอน save
	// t.Run("Send webhook success", func(t *testing.T) {
	// 	err := pkg.WebhookSend(validSyntaxText, envMock)
	// 	assert.Nil(t, err)
	// })

	t.Run("Send webhook fail: incorrect token", func(t *testing.T) {
		// mock incorrect token
		configs.Cfg.DiscordWebhook.Token = incorrectToken
		err := logs.WebhookSend(validSyntaxText, envMock)
		assert.Error(t, err)
		assert.ErrorIs(t, err, err.(*discordgo.RESTError))
		assert.Equal(t, 401, err.(*discordgo.RESTError).Response.StatusCode)
	})

	t.Run("Send webhook fail: id not found", func(t *testing.T) {
		// mock incorrect ID
		configs.Cfg.DiscordWebhook.ID = notFoundID
		err := logs.WebhookSend(validSyntaxText, envMock)
		assert.Error(t, err)
		assert.ErrorIs(t, err, err.(*discordgo.RESTError))
		assert.Equal(t, 404, err.(*discordgo.RESTError).Response.StatusCode)
	})

	t.Run("Send webhook fail: id invalid", func(t *testing.T) {
		// mock incorrect ID
		configs.Cfg.DiscordWebhook.ID = invalidID
		err := logs.WebhookSend(validSyntaxText, envMock)
		assert.Error(t, err)
		assert.ErrorIs(t, err, err.(*discordgo.RESTError))
		assert.Equal(t, 400, err.(*discordgo.RESTError).Response.StatusCode)
	})
}
