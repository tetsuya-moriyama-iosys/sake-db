package main

import (
    "context"
    "log"
    "os"
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	"backend/graph/model"
    "golang.org/x/crypto/bcrypt"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "github.com/99designs/gqlgen/graphql/handler"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "backend/graph/resolver"
    "backend/graph/generated"
    "github.com/joho/godotenv"
	"backend/router"
)

func Register(c *gin.Context) {
    var user model.User
    if err := c.BindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while hashing password"})
        return
    }
    user.Password = string(hashedPassword)
    user.ID = primitive.NewObjectID()

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    _, err = userCollection.InsertOne(ctx, user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating user"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

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
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
    if jwtSecretKey == "" {
        log.Fatal("JWT_SECRET_KEY environment variable is required")
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
    resolver := &resolver.Resolver{
		DB: db,
		SecretKey: jwtSecretKey,
	}

    srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	//ルーター作成
	r:=router.Router(srv)

   
    log.Println("connect to http://localhost:8080/ for GraphQL playground")
    log.Fatal(r.Run(":8080"))
}
