package main

import (
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type getBotsReturn struct {
	Ok             bool     `json:"ok"`
	ChangedKey     int      `json:"changed_key"`
	ChangedWebhook int      `json:"changed_webhook"`
	Working        int      `json:"working"`
	Result         []getBot `json"result"`
}

type getBot struct {
	Token            string `json:"token"`
	Username         string `json:"username"`
	Text             string `json:"text"`
	IsChangedKey     bool   `json:"is_changed_key"`
	IsWorking        bool   `json:"is_working"`
	IsChangedWebhook bool   `json:"is_changed_webhook"`
}

func (ses session) getBots(c *gin.Context) {

	keys, err := ses.Keys("*").Result()
	if err != nil {
		c.JSON(500, gin.H{
			"ok":          "false",
			"description": "Internal server error.",
		})
		return
	}

	c.JSON(200, ses.populateGetBotsReturn(keys))

}

func (ses session) populateGetBotsReturn(keys []string) getBotsReturn {

	var ret getBotsReturn
	ret.Ok = true

	for _, key := range keys {

		ck := len(ret.Result)
		token := ses.Get(key).String()

		ret.Result = append(ret.Result, getBot{
			Token: key,
			Text:  token,
		})

		bot, err := tgbotapi.NewBotAPI(key)
		if err != nil {
			ret.Result[ck].IsChangedWebhook = true
			ret.Result[ck].IsChangedKey = true
			ret.ChangedWebhook++
			ret.ChangedKey++
			continue
		}

		webhookInfo, _ := bot.GetWebhookInfo()
		if webhookInfo.URL != "https://"+ses.BackendHost+"/telegram/"+token {
			ret.Result[ck].IsChangedWebhook = true
			ret.ChangedWebhook++
			continue
		}

		ret.Working++
		ret.Result[ck].IsWorking = true

	}

	return ret

}
