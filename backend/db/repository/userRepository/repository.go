package userRepository

import (
	"backend/db"
	"context"
	"errors"
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
	result, err := r.collection.UpdateOne(ctx, bson.M{Email: email}, bson.M{
		"$set": bson.M{
			PasswordResetToken:       token,
			PasswordResetTokenExpire: time.Now().Add(1 * time.Hour),
		},
	})
	if result.MatchedCount == 0 {
		// ドキュメントが存在しなかった場合の処理
		return errors.New("ユーザーが見つかりません")
	}
	return err
}

func (r *UsersRepository) GetByPasswordToken(ctx context.Context, token string) (*Model, error) {
	// コレクションを取得
	var user Model
	if err := r.collection.FindOne(ctx, bson.M{PasswordResetToken: token, PasswordResetTokenExpire: bson.M{"$gt": time.Now()}}).Decode(&user); err != nil {
		return nil, errors.New("有効期限切れです。パスワードリセットURLを再発行してください。")
	}

	return &user, nil
}

func (r *UsersRepository) PasswordReset(ctx context.Context, user Model, password []byte) error {
	// パスワードを更新するクエリを実行
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": user.ID}, bson.M{
		"$set": bson.M{
			Password:                 password,
			PasswordResetToken:       nil,
			PasswordResetTokenExpire: nil,
		},
	})
	if err != nil {
		return err
	}
	return nil
}
