package routes

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "database/sql"
    "github.com/gin-contrib/sessions"
    "golang.org/x/crypto/bcrypt"
)

func HandleRoutes(router *gin.Engine, database *sql.DB) {
    router.GET("/", func(context *gin.Context) {
        session := sessions.Default(context)
        username := session.Get("username")

        context.HTML(http.StatusOK, "index.html", gin.H {
            "is_logged_in": username != nil,
            "username": username,
        })
    })

    router.GET("/register", func(context *gin.Context) {
        session := sessions.Default(context)
        username := session.Get("username")

        context.HTML(http.StatusOK, "register.html", gin.H {
            "is_logged_in": username != nil,
            "username": username,
        })
    })

    router.POST("/register", func(context *gin.Context) {
        username_input := context.PostForm("username-input")
        password_input := context.PostForm("password-input")

        hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password_input), bcrypt.DefaultCost)
        if err != nil {
            context.AbortWithError(http.StatusInternalServerError, err)
            return
        }

        var count int
        err = database.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", username_input).Scan(&count)
        if err != nil {
            context.AbortWithError(http.StatusInternalServerError, err)
            return
        }

        if count > 0 {
            context.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
            return
        }

        _, err = database.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username_input, hashedPassword)
        if err != nil {
            context.AbortWithError(http.StatusInternalServerError, err)
            return
        }

        context.Redirect(http.StatusSeeOther, "/login")
    })

    router.GET("/login", func(context *gin.Context) {
        session := sessions.Default(context)
        username := session.Get("username")

        context.HTML(http.StatusOK, "login.html", gin.H {
            "is_logged_in": username != nil,
            "username": username,
        })
    })

    router.POST("/login", func(context *gin.Context) {
        username_input := context.PostForm("username-input")
        password_input := context.PostForm("password-input")

        var stored_password string
        err := database.QueryRow("SELECT password FROM users WHERE username = ?", username_input).Scan(&stored_password)
        if err != nil {
            if err == sql.ErrNoRows {
                context.JSON(http.StatusUnauthorized, gin.H{"error": "Username does not exist"})
                return
            }
            context.AbortWithError(http.StatusInternalServerError, err)
            return
        }

        err = bcrypt.CompareHashAndPassword([]byte(stored_password), []byte(password_input))
        if err != nil {
            context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
            return
        }

        session := sessions.Default(context)
        session.Set("username", username_input)
        session.Save()

        context.Redirect(http.StatusSeeOther, "/")
    })

    router.POST("/logout", func(context *gin.Context) {
        session := sessions.Default(context)
        session.Clear()
        session.Save()

        context.SetCookie("user_session", "", -1, "/", "", false, true)

        context.Redirect(http.StatusSeeOther, "/")
    })
    
    router.GET("/profile/:username", func(context *gin.Context) {
        requested_username := context.Param("username")
        
        var count int
        err := database.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", requested_username).Scan(&count)
        if err != nil {
            context.AbortWithError(http.StatusInternalServerError, err)
            return
        }

        if count == 1 {
            session := sessions.Default(context)
            username := session.Get("username")

            context.HTML(http.StatusOK, "profile.html", gin.H {
                "requested_username": requested_username,
                "username": username,
                "is_logged_in": username != nil,
            })
        } else {
            context.String(http.StatusNotFound, "User not found")
        }

    })
}
