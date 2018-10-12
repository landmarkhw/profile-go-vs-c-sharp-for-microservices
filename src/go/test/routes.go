package status

import (
	"time"
	"github.com/gin-gonic/gin"
)

func getJsonSimple(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"success": true,
	})
}

var now = time.Now()

func getJsonComplex(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"id": 1,
		"items": []gin.H{
			{
				"id": 1,
				"another": "thing",
				"ts": now,
			},
			{
				"id": 2,
				"another": "thing",
				"ts": now,
			},
			{
				"id": 3,
				"another": "thing",
				"ts": now,
			},
			{
				"id": 4,
				"another": "thing",
				"ts": now,
			},
			{
				"id": 5,
				"another": "thing",
				"ts": now,
			},
			{
				"id": 6,
				"another": "thing",
				"ts": now,
			},
		},
		"success": true,
	})
}

// Defines routes used by the web server
func Routes(engine *gin.Engine) {
	engine.GET("/test/json/simple", getJsonSimple)
	engine.GET("/test/json/complex", getJsonSimple)
}
