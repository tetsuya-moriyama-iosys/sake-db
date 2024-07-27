package handlers

import (
    "context"
    "time"

    "github.com/graphql-go/graphql"
    "github.com/graphql-go/handler"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "backend/database"
    "backend/models"
)

// GraphQLスキーマの定義
var messageType = graphql.NewObject(graphql.ObjectConfig{
    Name: "Message",
    Fields: graphql.Fields{
        "id": &graphql.Field{
            Type: graphql.String,
        },
        "message": &graphql.Field{
            Type: graphql.String,
        },
    },
})

var queryType = graphql.NewObject(graphql.ObjectConfig{
    Name: "Query",
    Fields: graphql.Fields{
        "message": &graphql.Field{
            Type: messageType,
            Args: graphql.FieldConfigArgument{
                "id": &graphql.ArgumentConfig{
                    Type: graphql.String,
                },
            },
            Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                id := p.Args["id"].(string)
                objID, err := primitive.ObjectIDFromHex(id)
                if err != nil {
                    return nil, err
                }

                collection := database.GetCollection("messages")
                ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
                defer cancel()

                var message models.Message
                err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&message)
                if err != nil {
                    return nil, err
                }

                return message, nil
            },
        },
    },
})

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
    Query: queryType,
})

// GraphQLハンドラーの設定
func GraphQLHandler() *handler.Handler {
    return handler.New(&handler.Config{
        Schema:   &schema,
        Pretty:   true,
        GraphiQL: true, // GraphiQLインターフェースを有効にする
    })
}
