package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
    "context"
    "log"
    "backend/graph/generated"
    "backend/graph/model"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

// ID is the resolver for the id field.
func (r *messageResolver) ID(ctx context.Context, obj *model.Message) (string, error) {
    // 変更点: fmt.Errorf("not implemented: ID - id") を削除し、IDを文字列として返すように修正
    return obj.ID.Hex(), nil
}

// CreateMessage is the resolver for the createMessage field.
func (r *mutationResolver) CreateMessage(ctx context.Context, input model.NewMessage) (*model.Message, error) {
    message := &model.Message{
        ID:      primitive.NewObjectID(),
        Message: input.Message,
    }
    _, err := r.Collection.InsertOne(ctx, message)
    if err != nil {
        log.Printf("Error inserting message: %v", err) // エラーハンドリングの追加
        return nil, err
    }
    return message, nil
}

// Messages is the resolver for the messages field.
func (r *queryResolver) Messages(ctx context.Context) ([]*model.Message, error) {
    var messages []*model.Message
    cursor, err := r.Collection.Find(ctx, bson.M{})
    if err != nil {
        log.Printf("Error finding messages: %v", err) // エラーハンドリングの追加
        return nil, err
    }
    defer cursor.Close(ctx)
    for cursor.Next(ctx) {
        var message model.Message
        if err = cursor.Decode(&message); err != nil {
            log.Printf("Error decoding message: %v", err) // エラーハンドリングの追加
            return nil, err
        }
        messages = append(messages, &message)
    }
    return messages, nil
}

// Message returns generated.MessageResolver implementation.
func (r *Resolver) Message() generated.MessageResolver { return &messageResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type messageResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }