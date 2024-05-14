package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func M1(c *gin.Context) {
	fmt.Println("first middleware")
	c.Next()
}

func M2(c *gin.Context) {
	fmt.Println("second middleware")
	c.Next()
}

func M3(c *gin.Context) {
	fmt.Println("third middleware")
	c.Next()
}
