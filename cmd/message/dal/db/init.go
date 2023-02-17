package db

import (
	"context"
	"douyin/pkg/constants"
	"douyin/pkg/util"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const (
	DOUYIN    string = "douyin"
	messagedb string = "message"
	userdb    string = "user"
)

var client *mongo.Client
var messageCollection *mongo.Collection
var userCollection *mongo.Collection
var snowflake *util.SnowFlake

func Init() {
	clientOptions := options.Client().ApplyURI(constants.MongoDefaultDSN)
	// 连接 MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	messageCollection = client.Database(DOUYIN).Collection(messagedb)
	userCollection = client.Database(DOUYIN).Collection(userdb)
	snowflake = util.NewSnowFlake(constants.UserServiceMachineID)
}
