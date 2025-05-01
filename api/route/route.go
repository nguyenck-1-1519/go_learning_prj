package route

import (
	"time"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(timeout time.Duration, db *gorm.DB, gin *gin.Engine) {
	publicRouter := gin.Group("", logger.SetLogger())
	NewBookRouter(timeout, db, publicRouter)

	// protectedRouter := gin.Group("")
	// Middleware to verify AccessToken
	// protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	// All privates APIs
	// NewProfileRouter(timeout, db, protectedRouter)
	// NewTaskRouter(env, timeout, db, protectedRouter)
}
