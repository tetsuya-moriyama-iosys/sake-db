package main

import (
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "backend/database"
    "backend/handlers"
    "time"
)

func main() {
    database.ConnectDB()

    r := gin.Default()

    // CORSの設定
    config := cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"}, // フロントエンドのURLを指定
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }

    r.Use(cors.New(config))

    r.GET("/getById/:id", handlers.GetMessage)
    r.POST("/graphql", gin.WrapH(handlers.GraphQLHandler()))
     // GraphiQLインターフェースのためのGETリクエスト
     r.GET("/graphql", gin.WrapH(handlers.GraphQLHandler()))

    r.Run(":8080")
}
