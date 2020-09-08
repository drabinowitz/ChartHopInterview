package router

import (
	"github.com/gin-gonic/gin"
)

type route interface {
	Handle(c *gin.Context)
}
