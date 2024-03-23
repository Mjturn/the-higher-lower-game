package main

import (
    "log"
    "github.com/gin-gonic/gin"
    "github.com/Mjturn/the-higher-lower-game/backend/routes"
)

func main() {
    router := gin.Default()
    router.Static("/static", "../frontend/static")
    router.LoadHTMLGlob("../frontend/templates/*.html")
    routes.HandleRoutes(router)

    err := router.Run(":8080")
    if err != nil {
        log.Fatal(err)
    }
}
