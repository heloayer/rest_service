package interfaces

import (
	"github.com/heloayer/rest/internal/delivery"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.POST("/rest/substr/find", delivery.FindSubstring)

	r.POST("/rest/email/check", delivery.CheckEmail)

	r.POST("/rest/counter/add/:value", delivery.IncrementCounter)
	r.POST("/rest/counter/sub/:value", delivery.DecrementCounter)
	r.GET("/rest/counter/val", delivery.GetCounterValue)

	r.POST("/rest/hash/calc", delivery.CalcHash)
	r.GET("/rest/hash/result/:request_id", delivery.GetResult)

	r.POST("/rest/user", delivery.CreateUser)
	
	r.GET("/rest/user/:id", delivery.GetUser)
	r.PUT("/rest/user/:id", delivery.UpdateUser)
	r.DELETE("/rest/user/:id", delivery.DeleteUser)
}
