package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://vishnu:jiodhandhanadhan@cluster0.qzfkvf6.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

const dbName = "BudgetApplication"
const colName1 = "User"
const colName2 = "Budget"

//this is done so that we can access this accross the application
var UserCollection *mongo.Collection
var BudgetCollection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB Connection Successfull!!")

	UserCollection = client.Database(dbName).Collection(colName1)
	BudgetCollection = client.Database(dbName).Collection(colName2)

	fmt.Println("Collection Instance is Ready!!")
}
