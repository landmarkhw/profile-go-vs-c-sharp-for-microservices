package routes

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/landmarkhw/profile-go-vs-c-sharp-for-microservices/repositories"
)

func getPerson(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		panic(err)
	}
	person := personrepo.Get(id)
	ctx.JSON(200, gin.H{"data": person})
}

// Defines routes
func Person(engine *gin.Engine) {
	engine.GET("/person/:id", getPerson)
}
