package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var connectionString = getEnvValue("MONGODB_URI")
const dbName = "BudgetApplication"
const colName1 = "User"
const colName2 = "Budget"

//this is done so that we can access this accross the application
var UserCollection *mongo.Collection
var BudgetCollection *mongo.Collection

func init() {
	fmt.Println(connectionString)
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

//probably in a different file

//helper function
func getEnvValue(key string) string{
	err := godotenv.Load()
	if err != nil{
		log.Fatal(err)
	}

	return os.Getenv(key)
}