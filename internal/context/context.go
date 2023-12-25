package context

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AppContext struct {
	*gin.Context
	DB *gorm.DB
}
