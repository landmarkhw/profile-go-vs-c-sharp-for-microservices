package status

import (
	"github.com/gin-gonic/gin"
	"github.com/landmarkhw/profile-go-vs-c-sharp-for-microservices/database"
)

func getStatus(ctx *gin.Context) {
	dbSuccess := pfdb.Test()

	ctx.JSON(200, gin.H{
		"database": dbSuccess,
	})
}

// Defines routes used by the web server
func Routes(engine *gin.Engine) {
	engine.GET("/status", getStatus)
}
