package engine

import (
	"github.com/gin-gonic/gin"
	"github.com/landmarkhw/profile-go-vs-c-sharp-for-microservices/person"
	"github.com/landmarkhw/profile-go-vs-c-sharp-for-microservices/status"
	"github.com/landmarkhw/profile-go-vs-c-sharp-for-microservices/test"
)

// Run the REST API web server
func Run() {
	engine := gin.Default()

	person.Routes(engine)
	status.Routes(engine)
	test.Routes(engine)

	engine.Run(":5000")
}
