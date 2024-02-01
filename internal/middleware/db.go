package middleware

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Database middleware add database connection to context
func Database(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}

}
