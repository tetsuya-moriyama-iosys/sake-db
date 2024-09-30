package userRepository

import (
	"backend/db"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
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

func (r *UsersRepository) Register(ctx context.Context, user *Model) (*Model, error) {
	// MongoDBにデータを挿入(emailが重複していたらエラーを返すはず)
	result, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	// 挿入されたIDをuserのIDにセット
	user.ID = result.InsertedID.(primitive.ObjectID)

	return user, nil
}

func (r *UsersRepository) Update(ctx context.Context, user *Model) error {
	// 更新のためのフィルタを定義（IDで検索）
	filter := bson.M{"_id": user.ID}

	// MongoDBにデータを挿入
	re, err := r.collection.UpdateOne(ctx, filter, bson.M{
		"$set": user, // userオブジェクト内のフィールドをセット
	})
	log.Println(re)
	if err != nil {
		return err
	}

	return nil
}

func (r *UsersRepository) GetByEmail(ctx context.Context, email string) (*Model, error) {
	// ドキュメントを取得
	var user Model
	if err := r.collection.FindOne(ctx, bson.M{Email: email}).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UsersRepository) GetById(ctx context.Context, id primitive.ObjectID) (*Model, error) {
	// コレクションを取得
	var liquor Model
	if err := r.collection.FindOne(ctx, bson.M{Id: id}).Decode(&liquor); err != nil {
		if err == mongo.ErrNoDocuments {
			// ドキュメントが見つからない場合、nilを返す（エラーにはしない）
			return nil, nil
		}
		log.Println("デコードエラー:", err)
		return nil, err
	}

	return &liquor, nil
}

func (r *UsersRepository) SetPasswordToken(ctx context.Context, email string, token string) error {
	// ドキュメントの更新
	_, err := r.collection.UpdateOne(ctx, bson.M{Email: email}, bson.M{
		"$set": bson.M{
			PasswordResetToken:       token,
			PasswordResetTokenExpire: time.Now().Add(1 * time.Hour),
		},
	})
	return err
}
