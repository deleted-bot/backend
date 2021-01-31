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
	Total          int      `json:"total"`
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
	ret.Total = len(keys)

	for _, key := range keys {

		ck := len(ret.Result)
		text, _ := ses.Get(key).Result()

		ret.Result = append(ret.Result, getBot{
			Token: key,
			Text:  text,
		})

		bot, err := tgbotapi.NewBotAPI(key)
		if err != nil {
			ret.Result[ck].IsChangedWebhook = true
			ret.Result[ck].IsChangedKey = true
			ret.ChangedWebhook++
			ret.ChangedKey++
			continue
		}

		ret.Result[ck].Username = bot.Self.UserName

		webhookInfo, _ := bot.GetWebhookInfo()
		if webhookInfo.URL != "https://"+ses.BackendHost+"/telegram/"+key {
			ret.Result[ck].IsChangedWebhook = true
			ret.ChangedWebhook++
			continue
		}

		ret.Working++
		ret.Result[ck].IsWorking = true

	}

	return ret

}

func (ses session) unsetBot(c *gin.Context) {

	token := c.PostForm("token")

	doError := func(err error) {
		c.JSON(503, gin.H{
			"ok": true,
			"message": "Internal error",
			"details": err.Error(),
		})
	}

	err := ses.unsetWebhookForBot(token)
	if err != nil {
		doError(err)
		return
	}

	err = ses.Del(token).Err()
	if err != nil {
		doError(err)
		return
	}

	c.JSON(200, gin.H{
		"ok": true,
		"message": "Webhook removed",
		"details": "The webhook has been removed",
	})

}