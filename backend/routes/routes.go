package routes

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func HandleRoutes(router *gin.Engine) {
    router.GET("/", func(context *gin.Context) {
        context.HTML(http.StatusOK, "index.html", nil)
    })
    
    router.GET("/register", func(context *gin.Context) {
        context.HTML(http.StatusOK, "register.html", nil)
    })

    router.GET("/login", func(context *gin.Context) {
        context.HTML(http.StatusOK, "login.html", nil)
    })
}
