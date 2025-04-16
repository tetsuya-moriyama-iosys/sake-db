package liquorRepository

import (
	"backend/graph/graphModel"
	"backend/middlewares/customError"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	TagCollectionName = "liquors_tags"
)

type TagModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	LiquorId  primitive.ObjectID `bson:"liquor_id"`
	Text      string             `bson:"text"`
	UserId    primitive.ObjectID `bson:"user_id"`
	CreatedAt time.Time          `bson:"created_at"`
}

func (m *TagModel) ToGraphQL() *graphModel.Tag {
	return &graphModel.Tag{
		ID:   m.ID.Hex(),
		Text: m.Text,
	}
}

// TagsToGraphQL 複数のTagModelを変換
func TagsToGraphQL(tags []*TagModel) []*graphModel.Tag {
	var graphTags []*graphModel.Tag
	for _, tag := range tags {
		graphTags = append(graphTags, tag.ToGraphQL())
	}
	return graphTags
}

func (r *LiquorsRepository) GetTags(ctx context.Context, liquorId primitive.ObjectID) ([]*TagModel, *customError.Error) {
	cursor, err := r.tagCollection.Find(ctx, bson.M{LiquorID: liquorId})
	if err != nil {
		return nil, errGetTags(err, liquorId)
	}
	defer cursor.Close(ctx)

	// 結果を格納するスライス
	var tags []*TagModel
	if err = cursor.All(ctx, &tags); err != nil {
		return nil, errGetTagsDecode(err, liquorId)
	}

	return tags, nil
}

func (r *LiquorsRepository) PostTag(ctx context.Context, liquorId primitive.ObjectID, userId primitive.ObjectID, tag string) (*TagModel, *customError.Error) {
	newTag := &TagModel{
		LiquorId:  liquorId,
		Text:      tag,
		UserId:    userId,
		CreatedAt: time.Now(),
	}
	result, err := r.tagCollection.InsertOne(ctx, newTag)
	if err != nil {
		return nil, errPostTag(err, newTag)
	}

	newTag.ID = result.InsertedID.(primitive.ObjectID)
	return newTag, nil
}

func (r *LiquorsRepository) DeleteTag(ctx context.Context, id primitive.ObjectID) *customError.Error {
	result, err := r.tagCollection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return errDeleteTag(err, id)
	}

	if result.DeletedCount == 0 {
		return errZeroDelete(err, id)
	}

	return nil
}
