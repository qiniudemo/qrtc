package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

	if !CheckMac() {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Qrtc",
	})
}
