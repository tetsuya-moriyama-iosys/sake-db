package userRepository

import (
	"backend/db"
	"backend/middlewares/customError"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type UsersRepository struct {
	db         *db.DB
	collection *mongo.Collection //コレクションを先に取得して格納しておく
}

func NewUsersRepository(db *db.DB) UsersRepository {
	return UsersRepository{
		db:         db,
		collection: db.Collection(CollectionName),
	}
}

func (r *UsersRepository) Register(ctx context.Context, user *Model) (*Model, *customError.Error) {
	// MongoDBにデータを挿入
	result, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return nil, errRegister(err, user)
	}

	// 挿入されたIDをuserのIDにセット
	user.ID = result.InsertedID.(primitive.ObjectID)

	return user, nil
}

func (r *UsersRepository) Update(ctx context.Context, user *Model) *customError.Error {
	// MongoDBにデータを挿入
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": user.ID}, bson.M{
		"$set": user, // userオブジェクト内のフィールドをセット
	})
	if err != nil {
		return errUpdate(err, user)
	}

	return nil
}

func (r *UsersRepository) GetByEmail(ctx context.Context, email string) (*Model, *customError.Error) {
	// ドキュメントを取得
	var user Model
	if err := r.collection.FindOne(ctx, bson.M{Email: email}).Decode(&user); err != nil {
		return nil, errGetByEmail(err, email)
	}

	return &user, nil
}

func (r *UsersRepository) GetById(ctx context.Context, id primitive.ObjectID) (*Model, *customError.Error) {
	// コレクションを取得
	var user Model
	if err := r.collection.FindOne(ctx, bson.M{Id: id}).Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			// ドキュメントが見つからない場合、nilを返す（エラーにはしない）
			return nil, nil
		}
		return nil, errGetById(err, id)
	}

	return &user, nil
}

func (r *UsersRepository) GetByTwitterId(ctx context.Context, id string) (*Model, *customError.Error) {
	// コレクションを取得
	var user Model
	if err := r.collection.FindOne(ctx, bson.M{TwitterId: id}).Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			// ドキュメントが見つからない場合、nilを返す（エラーにはしない）
			return nil, nil
		}
		return nil, errGetByTwitterId(err, id)
	}
	return &user, nil
}

func (r *UsersRepository) SetPasswordToken(ctx context.Context, email string, token string) *customError.Error {
	// ドキュメントの更新
	result, err := r.collection.UpdateOne(ctx, bson.M{Email: email}, bson.M{
		"$set": bson.M{
			PasswordResetToken:       token,
			PasswordResetTokenExpire: time.Now().Add(1 * time.Hour),
		},
	})
	if result.MatchedCount == 0 {
		// ドキュメントが存在しなかった場合の処理
		return errSetPasswordTokenNotFound(email, token)
	}
	if err != nil {
		return errSetPasswordToken(err, email, token)
	}
	return nil
}

func (r *UsersRepository) GetByPasswordToken(ctx context.Context, token string) (*Model, *customError.Error) {
	// コレクションを取得
	var user Model
	if err := r.collection.FindOne(ctx, bson.M{PasswordResetToken: token, PasswordResetTokenExpire: bson.M{"$gt": time.Now()}}).Decode(&user); err != nil {
		return nil, errGetByPasswordToken(err, token)
	}

	return &user, nil
}

func (r *UsersRepository) PasswordReset(ctx context.Context, user Model, newPasswordHashed []byte) *customError.Error {
	// パスワードを更新するクエリを実行
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": user.ID}, bson.M{
		"$set": bson.M{
			Password:                 newPasswordHashed,
			PasswordResetToken:       nil,
			PasswordResetTokenExpire: nil,
		},
	})
	if err != nil {
		return errPasswordReset(err, user)
	}
	return nil
}
