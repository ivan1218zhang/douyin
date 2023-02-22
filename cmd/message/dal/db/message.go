package db

import (
	"context"
	"douyin/kitex_gen/message"
	"fmt"
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
	fmt.Println(from)
	fmt.Println(to)
	if from < to {
		index = strconv.FormatInt(to, 10) + "-" + strconv.FormatInt(from, 10)
	} else {
		index = strconv.FormatInt(from, 10) + "-" + strconv.FormatInt(to, 10)
	}
	var err error
	var cur *mongo.Cursor

	one := userCollection.FindOne(ctx, bson.M{"userid": to})
	fmt.Println(1)
	if one.Err() != nil {
		fmt.Println(2)
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
			fmt.Println(3)
			return res, err
		}
		fmt.Println(4)
	}
	fmt.Println(5)
	for cur.Next(ctx) {
		var tempMessage TempMessage
		fmt.Println(10)
		err := cur.Decode(&tempMessage)
		if err != nil {
			fmt.Println(6)
			return nil, err
		}
		res = append(res, &tempMessage)
		if cur.RemainingBatchLength() == 0 {
			var tempMessage1 []*TempMessage1
			err = cur.All(ctx, &tempMessage1)
			fmt.Println(7)

			if err != nil {
				fmt.Println(8)
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
	fmt.Println(res)
	return res, nil
}

// GetLatestMessage Get Message List
func GetLatestMessage(ctx context.Context, from int64, to int64) (message.Message, error) {
	var index string
	var res message.Message

	if from < to {
		index = strconv.FormatInt(to, 10) + "-" + strconv.FormatInt(from, 10)
	} else {
		index = strconv.FormatInt(from, 10) + "-" + strconv.FormatInt(to, 10)
	}
	var err error

	var cur *mongo.Cursor
	filter := bson.D{{"index", index}}
	cur, err = messageCollection.Find(ctx, filter, options.Find().SetSort(bson.D{{"_id", -1}}).SetLimit(1))
	if err != nil {
		return res, err
	}

	for cur.Next(ctx) {

		var tempMessage TempMessage

		err := cur.Decode(&tempMessage)
		if err != nil {
			return res, err
		}
		res.Id = tempMessage.ID
		res.CreateTime = tempMessage.Date
		res.FromUserId = tempMessage.FromUserID
		res.ToUserId = tempMessage.ToUserID
		res.Content = tempMessage.Content
	}

	return res, nil
}
