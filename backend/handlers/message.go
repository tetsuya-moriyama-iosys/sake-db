package handlers

import (
    "context"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "backend/database"
    "backend/models"
)

func GetMessage(c *gin.Context) {
    collection := database.GetCollection("messages")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    id := c.Param("id")
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var message models.Message
    err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&message)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, message)
}
