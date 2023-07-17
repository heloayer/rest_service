package delivery

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/heloayer/rest/initialize"
)

func IncrementCounter(c *gin.Context) {

	value, err := strconv.Atoi(c.Param("i"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Неверное значение"})
		return
	}

	err = initialize.Rdb.IncrBy("counter", int64(value)).Err()
	if err != nil {
		c.JSON(500, gin.H{"error": "Ошибка при увеличении счетчика"})
		return
	}

	c.JSON(200, gin.H{"message": "Счетчик увеличен"})
}

func DecrementCounter(c *gin.Context) {

	value, err := strconv.Atoi(c.Param("i"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Неверное значение"})
		return
	}

	err = initialize.Rdb.DecrBy("counter", int64(value)).Err()
	if err != nil {
		c.JSON(500, gin.H{"error": "Не получилось уменьшить счетчик"})
		return
	}

	c.JSON(200, gin.H{"message": "Счетчик уменьшен"})
}

func GetCounterValue(c *gin.Context) {

	val, err := initialize.Rdb.Get("counter").Result()
	if err != nil {
		c.JSON(500, gin.H{"error": "Ошибка получения значения счетчика"})
		return
	}

	value, err := strconv.Atoi(val)
	if err != nil {
		c.JSON(500, gin.H{"error": "Конвертация счетчика не удалась (to int)"})
		return
	}

	c.JSON(200, gin.H{"value": value})
}
