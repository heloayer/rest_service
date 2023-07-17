package delivery

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func FindSubstring(c *gin.Context) {
	input := c.Request.FormValue("input")
	if input == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "требуется входная строка"})
		return
	}

	maxSubstr := ""
	for i := 0; i < len(input); i++ {
		substr := ""
		for j := i; j < len(input); j++ {
			if strings.Contains(substr, string(input[j])) {
				break
			}
			substr += string(input[j])
		}
		if len(substr) > len(maxSubstr) {
			maxSubstr = substr
		}
	}

	c.JSON(http.StatusOK, gin.H{"maxSubstring": maxSubstr})
}
