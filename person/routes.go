package person

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func getPerson(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		panic(err)
	}
	person := Get(id)
	ctx.JSON(200, gin.H{"data": person})
}

// Defines routes
func Routes(engine *gin.Engine) {
	engine.GET("/person/:id", getPerson)
}
