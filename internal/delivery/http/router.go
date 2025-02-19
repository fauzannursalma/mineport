package http

import (
	"github.com/fauzannursalma/mineport/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()

    authGroup := router.Group("/auth")
    {
        authGroup.POST("/register", RegisterHandler)
        authGroup.POST("/login", LoginHandler)
    }

    // Protected routes
    apiGroup := router.Group("/api")
    apiGroup.Use(middleware.AuthMiddleware()) 
    {
        apiGroup.GET("/users/me", GetUserProfile)
        apiGroup.GET("/skills", GetAllSkills)
        apiGroup.GET("/projects", GetAllProjects)
    }

    return router
}
