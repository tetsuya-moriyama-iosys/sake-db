package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"backend/graph/model"
	"context"
)

// Liquor is the resolver for the liquor field.
func (r *queryResolver) Liquor(ctx context.Context, id string) (*model.Liquor, error) {
	// idをObjectIDに変換
	//objectID, err := primitive.ObjectIDFromHex(id)
	//if err != nil {
	//	return nil, fmt.Errorf("無効なID形式: %s", id)
	//}
	//
	//// コレクションを取得
	//var liquor liquorModel.Schema
	//if err := db.GetCollection(liquorModel.CollectionName).FindOne(context.TODO(), bson.M{liquorModel.ID: objectID}).Decode(&liquor); err != nil {
	//	log.Println("デコードエラー:", err)
	//	return nil, err
	//}
	//return liquor.ToGraphQL(), nil
	return nil, nil
}

// RandomRecommendList is the resolver for the randomRecommendList field.
func (r *queryResolver) RandomRecommendList(ctx context.Context, limit int) ([]*model.Liquor, error) {
	// コレクションを取得
	//collection := db.GetCollection(liquorModel.CollectionName)
	//
	//// $sampleパイプラインを使ってランダムに指定件数を取得
	//cursor, err := collection.Aggregate(ctx, mongo.Pipeline{
	//	{{"$sample", bson.D{{"size", limit}}}},
	//})
	//if err != nil {
	//	return nil, err
	//}
	//defer cursor.Close(ctx)
	//
	//var collections []*liquorModel.Schema
	//if err = cursor.All(ctx, &collections); err != nil {
	//	return nil, err
	//}
	//
	//var results []*model.Liquor
	//// 結果をGraphQLの構造体にマッピング（スネークケースからキャメルケースへの変換）
	//for _, liquor := range collections {
	//	results = append(results, liquor.ToGraphQL())
	//}
	//
	//return results, nil

	return nil, nil
}
