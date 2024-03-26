package main

import (
    "fmt"
    "log"
    "os"
    "github.com/gin-gonic/gin"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "github.com/joho/godotenv"
    "github.com/gin-contrib/sessions"
    "github.com/gin-contrib/sessions/cookie"
    "github.com/Mjturn/the-higher-lower-game/backend/routes"
)

func main() {
    err := godotenv.Load("../.env")
    if err != nil {
        log.Fatal(err)
    }

    database_username := os.Getenv("DATABASE_USERNAME")
    database_password := os.Getenv("DATABASE_PASSWORD")
    database_host := os.Getenv("DATABASE_HOST")
    database_port := os.Getenv("DATABASE_PORT")
    database_name := os.Getenv("DATABASE_NAME")

    database_connection_string := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", database_username, database_password, database_host, database_port, database_name)

    database, err := sql.Open("mysql", database_connection_string)
    if err != nil {
        log.Fatal(err)
    }
    defer database.Close()

    err = database.Ping()
    if err != nil {
        log.Fatal(err)
    }
    
    router := gin.Default()
    router.Static("/static", "../frontend/static")
    router.LoadHTMLGlob("../frontend/templates/*.html")
    store := cookie.NewStore([]byte(os.Getenv("SESSION_SECRET_KEY")))
    router.Use(sessions.Sessions("user_session", store))
    routes.HandleRoutes(router, database)

    err = router.Run(":8080")
    if err != nil {
        log.Fatal(err)
    }
}
