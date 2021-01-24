package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"encoding/json"
	"io/ioutil"
)

func (ses session) setWebhookForBot(token string) (string, error) {

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return "", err
	}

	_, err = bot.Request(tgbotapi.NewWebhook("https://" + ses.Host + "/telegram/" + token))
	if err != nil {
		return "", err
	}

	return bot.Self.UserName, nil

}

func (ses session) telegram(c *gin.Context) {

	token := c.Param("token")
	bot, err1 := tgbotapi.NewBotAPI(token)
	body, err2 := ioutil.ReadAll(c.Request.Body)

	if err1 != nil || err2 != nil {
		return
	}

	var update tgbotapi.Update
	json.Unmarshal(body, &update)

	if update.Message != nil {

		if update.Message.Chat.ID > 0 {

			val, err := ses.Get(token).Result()
			if err != nil {
				return
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, val)
			msg.ParseMode = "HTML"
			bot.Send(msg)

			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Made with " + ses.Host))

		}

	}

}
