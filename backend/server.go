package main

import (
    "context"
    "log"
    "os"

    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "github.com/99designs/gqlgen/graphql/handler"
    "github.com/99designs/gqlgen/graphql/playground"
    "backend/graph"
    "backend/graph/generated"
    "github.com/joho/godotenv"
)

func main() {
    // .envファイルを読み込みます
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // 環境変数からMongoDB URIを取得します
    mongoURI := os.Getenv("MONGO_URI")
    if mongoURI == "" {
        log.Fatal("MONGO_URI environment variable is required")
    }

    clientOptions := options.Client().ApplyURI(mongoURI)

    client, err := mongo.NewClient(clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    ctx := context.TODO()
    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }

    collection := client.Database("helloworld").Collection("messages")

    resolver := &graph.Resolver{Collection: collection}

    srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

    r := gin.Default()
    r.POST("/query", func(c *gin.Context) {
        srv.ServeHTTP(c.Writer, c.Request)
    })
    r.GET("/", func(c *gin.Context) {
        playground.Handler("GraphQL", "/query").ServeHTTP(c.Writer, c.Request)
    })

    log.Println("connect to http://localhost:8080/ for GraphQL playground")
    log.Fatal(r.Run(":8080"))
}
