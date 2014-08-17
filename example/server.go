package main

import (
	"github.com/gin-gonic/gin"
	"github.com/k0kubun/fluent-logger-go"
)

var (
	logger *fluent.Logger
)

func logId(c *gin.Context) {
	id := c.Params.ByName("id")
	logger.Post("logId", map[string]string{"id": id})
	c.String(200, "OK")
}

func main() {
	logger = fluent.NewLogger(fluent.Config{
		BufferLength: 3 * 1024,
	})

	r := gin.Default()
	r.GET("/:id", logId)
	r.Run(":3000")
}
