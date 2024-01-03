package dbagents

import (
	"context"
	"fmt"
	"log"
	helpers "mongoDB/Helpers"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Agents *mongo.Collection
var Client *mongo.Client
 




func DBconnectionAgents(){
godotenv.Load()
dbConnection := options.Client().ApplyURI(os.Getenv("connectionString"))
actualConnection,err := mongo.Connect(context.TODO(),dbConnection)

if err != nil {
panic(err.Error())	

}


Client= actualConnection

Agents = Client.Database(os.Getenv("databaseName")).Collection(os.Getenv("agentCollection"))
fmt.Println("Connected to DB Agents")

//make agentnumber uniqe 
agentNumberOption := options.Index().SetUnique(true)
agentNumberindex := mongo.IndexModel{ 
	Keys: bson.M{"agentNo":1},
	Options: agentNumberOption,
}
 
_, err1 := Agents.Indexes().CreateOne(context.Background(),agentNumberindex)
if err1 != nil {
	log.Fatal(err1.Error())
	return
}



Storenumberoption := options.Index().SetUnique(true)
StoreNumberindex := mongo.IndexModel{ 
	Keys: bson.M{"storeNo":1},
	Options: Storenumberoption,
}

_, err2 := Agents.Indexes().CreateOne(context.Background(),StoreNumberindex)
if err2 != nil {
	log.Fatal(err2.Error())
	return
}



Agentnameoption := options.Index().SetUnique(true)
Agentnameindex := mongo.IndexModel{ 
	Keys: bson.M{"agentname":1},
	Options: Agentnameoption,
}

_, err3 := Agents.Indexes().CreateOne(context.Background(),Agentnameindex)
if err3 != nil {
	log.Fatal(err3.Error())
	return
}




helpers.Wg.Done()
}
