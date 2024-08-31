package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"backend/graph/generated"
	"backend/graph/model"
	"context"
)

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, username string, email string, password string) (*model.User, error) {
	//hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	//if err != nil {
	//	return nil, err
	//}
	//
	//now := time.Now()
	//user := &model.User{
	//	Username:  username,
	//	Email:     email,
	//	Password:  string(hashedPassword),
	//	CreatedAt: now,
	//	UpdatedAt: now,
	//}
	//
	//_, err = r.usersCollection().InsertOne(ctx, user)
	//
	//if err != nil {
	//	log.Printf("Error inserting user: %v", err) // エラーハンドリングの追加
	//	return nil, err
	//}
	//
	//return user, nil
	return nil, nil
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, username string, password string) (string, error) {
	//var user model.User
	//
	//err := r.usersCollection().FindOne(ctx, bson.M{"username": username}).Decode(&user)
	//if err != nil {
	//	return "", errors.New("user not found")
	//}
	//
	//err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	//if err != nil {
	//	return "", errors.New("invalid password")
	//}
	//
	//token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	//	"username": user.Username,
	//	"exp":      time.Now().Add(time.Hour * 72).Unix(),
	//})
	//
	//tokenString, err := token.SignedString([]byte(r.SecretKey))
	//if err != nil {
	//	return "", err
	//}
	//
	//return tokenString, nil

	return "", nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	//var users []*model.User
	//
	//cursor, err := r.usersCollection().Find(ctx, bson.M{})
	//if err != nil {
	//	return nil, err
	//}
	//defer cursor.Close(ctx)
	//for cursor.Next(ctx) {
	//	var user model.User
	//	if err := cursor.Decode(&user); err != nil {
	//		return nil, err
	//	}
	//	users = append(users, &user)
	//}
	//return users, nil

	return nil, nil
}

// ID is the resolver for the id field.
func (r *userResolver) ID(ctx context.Context, obj *model.User) (string, error) {
	return obj.ID.Hex(), nil
}

// CreatedAt is the resolver for the createdAt field.
func (r *userResolver) CreatedAt(ctx context.Context, obj *model.User) (string, error) {
	return obj.CreatedAt.Format("2006-01-02 15:04:05"), nil
}

// UpdatedAt is the resolver for the updatedAt field.
func (r *userResolver) UpdatedAt(ctx context.Context, obj *model.User) (string, error) {
	return obj.UpdatedAt.Format("2006-01-02 15:04:05"), nil
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
