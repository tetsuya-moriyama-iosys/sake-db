package main

import (
    "context"
    "log"
	"fmt"
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
    "go.mongodb.org/mongo-driver/bson"
    "backend/graph/resolver"
    "backend/graph/generated"
    "github.com/joho/godotenv"
	"backend/router"
)

func Register(c *gin.Context, db *mongo.Database) {
    // フロントエンドからのデータを受け取るための一時構造体
    var requestData struct {
        Name     string `json:"name"`
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    if err := c.BindJSON(&requestData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // デバッグ用に受信データをログ出力
	log.Println("Received user data: %+v\n", requestData)

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(requestData.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while hashing password"})
        return
    }

    // User構造体にマッピング
    user := model.User{
        ID:       primitive.NewObjectID(),
        Username:     requestData.Name, // データベースのフィールド名は `Name`
        Email:    requestData.Email,
        Password: string(hashedPassword),
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    _, err = db.Collection("users").InsertOne(ctx, user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating user"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully!!!!!!!!!!"})
}

func Login(c *gin.Context, db *mongo.Database) {
    var user model.User
    var foundUser model.User

    if err := c.BindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err := db.Collection("users").FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
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

	// ユーザー登録とログインのエンドポイント
    r.POST("/register", func(c *gin.Context) {
        Register(c, db)
    })
    r.POST("/login", func(c *gin.Context) {
        Login(c, db)
    })

   
    log.Println("connect to http://localhost:8080/ for GraphQL playground")
    log.Fatal(r.Run(":8080"))
}
