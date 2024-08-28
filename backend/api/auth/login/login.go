package login

import (
	"backend/graph/model"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	"backend/db"
	"go.mongodb.org/mongo-driver/bson"
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
