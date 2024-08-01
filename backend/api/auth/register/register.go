package register

import(
	"time"
    "context"
	"backend/graph/model"
    "log"
	"net/http"
	"github.com/gin-gonic/gin"

    "go.mongodb.org/mongo-driver/mongo"
    "golang.org/x/crypto/bcrypt"
    "go.mongodb.org/mongo-driver/bson/primitive"
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