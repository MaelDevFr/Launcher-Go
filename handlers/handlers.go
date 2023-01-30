package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeFunc(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{
		"title": "Home",
	})
}