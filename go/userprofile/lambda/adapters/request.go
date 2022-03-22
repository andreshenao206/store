package adapters

import (
	"context"
	"fmt"
	"log"

	"github.com/store/go/userprofile/lambda/config"
	"github.com/store/go/userprofile/lambda/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var appConfig = config.LoadConfig()

func RequestNode(req domain.User, c string) (str string, err error) {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(appConfig.NetworkHost))
	if err != nil {
		log.Printf("::::POST:::: mongo.Connect %v", err)
		return
	}

	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Printf("::::POST:::: client.Ping %v", err)
		return
	}

	usersCollection := client.Database(appConfig.NetworkDB).Collection(c)

	user := bson.D{
		primitive.E{Key: "fullname", Value: req.FullName}, {Key: "age", Value: req.Age},
	}
	result, err := usersCollection.InsertOne(context.TODO(), user)

	if err != nil {
		log.Printf("::::POST:::: usersCollection.InsertOne %v", err)
		return
	}

	str = fmt.Sprintf("%v", result)
	return

}
