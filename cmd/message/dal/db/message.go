package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	bson2 "gopkg.in/mgo.v2/bson"
	"strconv"
	"time"
)

type TempMessage struct {
	ID         int64
	FromUserID int64
	ToUserID   int64
	Content    string
	Date       int64
	Index      string
}

type TempMessage1 struct {
	ObjectID   bson2.ObjectId `bson:"_id" `
	ID         int64
	FromUserID int64
	ToUserID   int64
	Content    string
	Date       int64
	Index      string
}
type user struct {
	UserID   int64
	ObjectID string `bson:"objectid" `
}

// SendMessage send message
func SendMessage(ctx context.Context, from int64, to int64, content string) error {
	var index string
	if from < to {
		index = strconv.FormatInt(to, 10) + "-" + strconv.FormatInt(from, 10)
	} else {
		index = strconv.FormatInt(from, 10) + "-" + strconv.FormatInt(to, 10)
	}
	tempMessage := TempMessage{snowflake.NextSnowID(), from, to, content, time.Now().Unix(), index}

	_, err := messageCollection.InsertOne(ctx, tempMessage)
	if err != nil {
		return err
	}

	return nil
}

// GetMessageList Get Message List
func GetMessageList(ctx context.Context, from int64, to int64) ([]*TempMessage, error) {
	var index string
	var res []*TempMessage

	if from < to {
		index = strconv.FormatInt(to, 10) + "-" + strconv.FormatInt(from, 10)
	} else {
		index = strconv.FormatInt(from, 10) + "-" + strconv.FormatInt(to, 10)
	}
	var err error
	var cur *mongo.Cursor

	one := userCollection.FindOne(ctx, bson.M{"userid": to})
	if one.Err() != nil {

		cur, err = messageCollection.Find(ctx, bson.M{"index": index}, options.Find().SetSort(bson.D{{"_id", 1}}))
		if err != nil {
			return res, err
		}
	} else {

		var tempUser user
		err = one.Decode(&tempUser)
		if err != nil {
			return res, err
		}

		objID, _ := primitive.ObjectIDFromHex(tempUser.ObjectID)

		filter := bson.D{{"fromuserid", from}, {"touserid", to}, {"_id", bson.M{"$gt": objID}}}
		cur, err = messageCollection.Find(ctx, filter, options.Find().SetSort(bson.D{{"_id", 1}}))
		if err != nil {
			return res, err
		}
	}

	for cur.Next(ctx) {
		var tempMessage TempMessage
		err := cur.Decode(&tempMessage)
		if err != nil {
			return nil, err
		}
		res = append(res, &tempMessage)
		if cur.RemainingBatchLength() == 0 {
			var tempMessage1 []*TempMessage1
			err = cur.All(ctx, &tempMessage1)
			if err != nil {
				return nil, err
			}
			opts := options.Update().SetUpsert(true)
			filter := bson.M{"userid": to}

			update := bson.D{{"$set", bson.D{{"objectid", tempMessage1[len(tempMessage1)-1].ObjectID}}}}

			_, err = userCollection.UpdateOne(ctx, filter, update, opts)
			if err != nil {
				return nil, err
			}
		}
	}

	return res, nil
}
