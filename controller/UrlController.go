package controller

import (
	"net/http"
	"urlshortner/constant"
	"urlshortner/types"

	"github.com/gin-gonic/gin"
)

func ShortTheUrl(c *gin.Context) {
	var shortUrlBody types.ShortUrlBody
	err := c.BindJSON(shortUrlBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": constant.BindError})
		return
	}
}
