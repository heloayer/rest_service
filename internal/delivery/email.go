package delivery

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

func CheckEmail(c *gin.Context) {
	body, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	emailRgx := regexp.MustCompile(`\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}\b`)
	emails := emailRgx.FindAllString(string(body), -1)

	c.JSON(http.StatusOK, gin.H{"emails": emails})
}
