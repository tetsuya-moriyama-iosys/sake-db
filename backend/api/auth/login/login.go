package login

import(
	"time"
    "context"
	"backend/graph/model"
	"net/http"
	"github.com/gin-gonic/gin"

    "golang.org/x/crypto/bcrypt"
	
    "go.mongodb.org/mongo-driver/bson"
	"backend/db"
)

func Login(c *gin.Context) {
	//データベースに接続
	db.ConnectDB()
	
    var user model.User
    var foundUser model.User

    if err := c.BindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err := db.GetCollection("users").FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
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