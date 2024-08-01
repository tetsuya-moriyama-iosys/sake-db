package main

import (
    "context"
    "log"
    "os"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "github.com/99designs/gqlgen/graphql/handler"
    "backend/graph/resolver"
    "backend/graph/generated"
    "github.com/joho/godotenv"
	"backend/router"
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

	//ハローワールドDBへのインスタンスを生成
    db := client.Database("helloworld")

	//リゾルバを設定
    resolver := &resolver.Resolver{DB: db}

    srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	//ルーター作成
	r:=router.Router(srv)

   
    log.Println("connect to http://localhost:8080/ for GraphQL playground")
    log.Fatal(r.Run(":8080"))
}
