package db

import (
	"context"
	"fmt"
	"log"
	helpers "mongoDB/Helpers"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Users *mongo.Collection
var Client *mongo.Client

func DBconnection() {
	dbConnection := options.Client().ApplyURI(os.Getenv("connectionString"))
	actualConnection, err := mongo.Connect(context.TODO(), dbConnection)

	if err != nil {
		panic(err.Error())

	}

	Client = actualConnection

	Users = Client.Database(os.Getenv("databaseName")).Collection(os.Getenv("usersCollection"))
	fmt.Println("Connected to DB Channels")

	//make idNo uniqe
	idNumberOption := options.Index().SetUnique(true)
	idNumberindex := mongo.IndexModel{
		Keys:    bson.M{"idno": 1},
		Options: idNumberOption,
	}

	_, err1 := Users.Indexes().CreateOne(context.Background(), idNumberindex)
	if err1 != nil {
		log.Fatal(err1.Error())
		return
	}

	PhoneOption := options.Index().SetUnique(true)
	Phoneindex := mongo.IndexModel{
		Keys:    bson.M{"phone": 1},
		Options: PhoneOption,
	}

	_, err2 := Users.Indexes().CreateOne(context.Background(), Phoneindex)
	if err2 != nil {
		log.Fatal(err2.Error())
		return
	}

	//close the wait grp

	helpers.Wg.Done()
}
