package engine

import (
	"github.com/gin-gonic/gin"
	"github.com/landmarkhw/profile-go-vs-c-sharp-for-microservices/routes"
)

// Run the REST API web server
func Run() {
	engine := gin.Default()

	routes.Status(engine)
	routes.Person(engine)

	engine.Run(":5000")
}
