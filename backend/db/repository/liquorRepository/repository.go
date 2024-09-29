package liquorRepository

import (
	"backend/db"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math/rand"
	"time"
)

type LiquorsRepository struct {
	db              *db.DB
	collection      *mongo.Collection //コレクションを先に取得して格納しておく
	logsCollection  *mongo.Collection
	boardCollection *mongo.Collection
}

func NewLiquorsRepository(db *db.DB) LiquorsRepository {
	return LiquorsRepository{
		db:              db,
		collection:      db.Collection(CollectionName),
		logsCollection:  db.Collection(LogsCollectionName),
		boardCollection: db.Collection(BoardCollectionName),
	}
}

func (r *LiquorsRepository) GetLiquorById(ctx context.Context, id primitive.ObjectID) (*Model, error) {
	// コレクションを取得
	var liquor Model
	if err := r.collection.FindOne(ctx, bson.M{ID: id}).Decode(&liquor); err != nil {
		return nil, err
	}

	return &liquor, nil
}
func (r *LiquorsRepository) GetLiquorByName(ctx context.Context, name string) (*Model, error) {
	// コレクションを取得
	var liquor Model
	if err := r.collection.FindOne(ctx, bson.M{Name: name}).Decode(&liquor); err != nil {
		return nil, err
	}

	return &liquor, nil
}
func (r *LiquorsRepository) GetLiquorByRandomKey(ctx context.Context, key float64) (*Model, error) {
	// コレクションを取得
	var liquor Model
	if err := r.collection.FindOne(ctx, bson.M{RandomKey: key}).Decode(&liquor); err != nil {
		return nil, err
	}

	return &liquor, nil
}

func (r *LiquorsRepository) GetLiquorsByIds(ctx context.Context, ids []primitive.ObjectID) ([]Model, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	// コレクションからフィルタに一致するドキュメントを取得
	cursor, err := r.collection.Find(ctx, bson.M{ID: bson.M{"$in": ids}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// 結果を格納するスライス
	var liquors []Model

	// 取得したドキュメントをスライスにデコード
	if err = cursor.All(ctx, &liquors); err != nil {
		return nil, err
	}

	return liquors, nil
}

func (r *LiquorsRepository) GetRandomLiquors(ctx context.Context, limit int) ([]*Model, error) {
	var collections []*Model

	// コレクションの総ドキュメント数を取得
	count, err := r.collection.CountDocuments(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, nil // ドキュメントがない場合
	}

	//100万件以下であれば、普通にsampleで取得する
	if count < 1000000 {
		// $sampleパイプラインを使ってランダムに指定件数を取得
		cursor, err := r.collection.Aggregate(ctx, mongo.Pipeline{
			{{"$sample", bson.D{{"size", limit}}}},
		})
		if err != nil {
			return nil, err
		}
		defer cursor.Close(ctx)

		if err = cursor.All(ctx, &collections); err != nil {
			return nil, err
		}

		return collections, nil
	}

	//100万件以上ある場合、ランダムキーによる取得ロジックにする
	randomValue := rand.New(rand.NewSource(time.Now().UnixNano())).Float64() // ランダムな基準値を生成
	// randomKey がランダム値以上のドキュメントを取得
	cursor, err := r.collection.Find(ctx, bson.M{
		RandomKey: bson.M{"$gte": randomValue},
	}, options.Find().SetLimit(int64(limit)))

	//noDocumentsエラーはFindメソッドでは返されない
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &collections); err != nil {
		return nil, err
	}

	// 取得した配列の長さが足らない場合、逆方向で探索
	if len(collections) < limit {
		remaining := limit - len(collections) //取得すべき件数

		// 残りを補完するために randomKey < randomValue の範囲で再クエリ
		cursor, err = r.collection.Find(ctx, bson.M{
			"randomKey": bson.M{"$lt": randomValue},
		}, options.Find().SetLimit(int64(remaining)))

		if err != nil {
			return nil, err // 再クエリでエラーが発生した場合は終了
		}

		var moreResults []*Model
		if err := cursor.All(ctx, &moreResults); err != nil {
			return nil, err // 再クエリのカーソル操作エラー
		}

		collections = append(collections, moreResults...)
	}

	// 取得した結果が足りなくてもエラーにせず、空の場合もそのまま返す
	return collections, nil

}

func (r *LiquorsRepository) GetLiquorsFromCategoryIds(ctx context.Context, ids []int) ([]*Model, error) {
	// クエリフィルターを作成。カテゴリIDがidsのいずれかに一致するリカーを取得
	filter := bson.M{"category_id": bson.M{"$in": ids}}

	// コレクションからフィルタに一致するドキュメントを取得
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// 結果を格納するスライス
	var liquors []*Model

	// 取得したドキュメントをスライスにデコード
	if err = cursor.All(ctx, &liquors); err != nil {
		return nil, err
	}

	return liquors, nil
}

func (r *LiquorsRepository) InsertOne(ctx context.Context, liquor *Model) (primitive.ObjectID, error) {
	result, err := r.collection.InsertOne(ctx, liquor)
	if err != nil {
		return primitive.NilObjectID, err
	}

	// InsertOneResultからIDを取得
	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, err
	}

	return id, nil
}

func (r *LiquorsRepository) UpdateOne(ctx context.Context, liquor *Model) (primitive.ObjectID, error) {
	// フィルタ：IDを用いてドキュメントを特定
	filter := bson.M{"_id": liquor.ID}

	// 構造体を BSON にマッピング
	data, err := bson.Marshal(liquor)
	if err != nil {
		return primitive.NilObjectID, err
	}

	// BSON を bson.M に変換
	var update bson.M
	if err := bson.Unmarshal(data, &update); err != nil {
		return primitive.NilObjectID, err
	}

	// 更新内容：$setオペレーターを使って指定したフィールドを更新
	updateBson := bson.M{"$set": update}

	// UpdateOneでドキュメントを更新
	result, err := r.collection.UpdateOne(ctx, filter, updateBson)
	if err != nil {
		return primitive.NilObjectID, err
	}

	// UpdateOneは更新したドキュメントのIDを直接返さないため、元のIDを返す
	if result.MatchedCount == 0 {
		return primitive.NilObjectID, fmt.Errorf("no document matched the provided ID")
	}

	return liquor.ID, nil
}

// UpdateRate 掲示板のratesを更新する
func (r *LiquorsRepository) UpdateRate(ctx context.Context, lId primitive.ObjectID, userId primitive.ObjectID, rate *int) error {
	// フィルタ：IDを用いてドキュメントを特定
	filter := bson.M{"_id": lId}
	// 以前の評価を全てのrate配列から削除（このユーザーの評価は一つだけ存在する想定）
	pullUpdate := bson.M{
		"$pull": bson.M{
			Rate5Users: userId,
			Rate4Users: userId,
			Rate3Users: userId,
			Rate2Users: userId,
			Rate1Users: userId,
		},
	}
	// pull操作で、過去の評価を削除
	_, err := r.collection.UpdateOne(ctx, filter, pullUpdate)
	if err != nil {
		return err
	}

	//未評価の場合は単純に評価を消して終わり
	if rate == nil {
		return nil
	}

	// 新しい評価をaddToSetで追加（重複しないようにする）
	rateField := fmt.Sprintf("rate%d_users", *rate)
	addToSetUpdate := bson.M{
		"$addToSet": bson.M{
			rateField: userId, // 新しい評価を追加
		},
	}
	// addToSet操作で評価を更新
	_, err = r.collection.UpdateOne(ctx, filter, addToSetUpdate)
	if err != nil {
		return err
	}
	return nil
}
