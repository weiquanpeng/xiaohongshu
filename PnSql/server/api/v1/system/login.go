package system

import "github.com/gin-gonic/gin"

type Login struct{}

func (b *Login) Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 200,
	})
}
