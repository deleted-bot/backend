package main

import (
	"github.com/gin-gonic/gin"
)

func (ses session) setBot(c *gin.Context) {

	token := c.PostForm("token")
	text := c.PostForm("text")

	if token == "" || text == "" || c.PostForm("tec") != "yes" {
		c.JSON(403, gin.H{
			"ok":      false,
			"message": "Missing fields",
			"details": "Please fill all the form data.",
		})
		return
	}

	botUsername, err := ses.setWebhookForBot(token)
	if err != nil {
		c.JSON(300, gin.H{
			"ok":      false,
			"message": "Bad token",
			"details": err.Error(),
		})
		return
	}

	err = ses.Set(token, text, 0).Err()

	if err != nil {
		c.JSON(503, gin.H{
			"ok":      false,
			"message": "Internal server error, please try again",
			"details": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"ok":      true,
		"message": "Created succesfully!",
		"details": "Your bot should be working now. You can check it out at <a href='https://t.me/" + botUsername + "'>@" + botUsername + "</a>",
	})

}
