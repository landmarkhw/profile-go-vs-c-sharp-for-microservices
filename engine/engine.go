package engine

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"landmarkhw.com/repositories"
)

func getStatus(ctx *gin.Context) {
	dbSuccess := personrepo.Test()

	ctx.JSON(200, gin.H{
		"database": dbSuccess,
	})
}

func getPerson(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		panic(err)
	}
	person := personrepo.Get(id)
	ctx.JSON(200, gin.H{"data": person})
}

// Defines routes used by the web server
func defineRoutes(engine *gin.Engine) {
	engine.GET("/status", getStatus)
	engine.GET("/person/:id", getPerson)
}

// Run the REST API web server
func Run() {
	engine := gin.Default()
	defineRoutes(engine)
	engine.Run(":5000")
}
