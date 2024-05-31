package database

import (
	"context"
	"fmt"
	"log"

	"github.com/vishnumenon/budgetapplication/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AddUser(user models.User) {

	inserted, err := UserCollection.InsertOne(context.Background(), user)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("User has been inserted Successfully with ID: ", inserted.InsertedID)
}

func UpdateUser(user models.User, movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{
		"name":     user.Name,
		"email":    user.Email,
		"password": user.Password,
		"role":     user.Role,
	}}

	updatedUser, err := UserCollection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("User Updated with object ID: ", movieId, " no: ", updatedUser.ModifiedCount)

}

func CheckUser(email string) bool {

	var result models.User

	filter := bson.M{"email": email}
	project := bson.M{"password": 0}
	opts := options.FindOne().SetProjection(project)

	err := UserCollection.FindOne(context.Background(), filter, opts).Decode(&result)

	if err != nil {
		log.Fatal(err)
	}

	//if result is empty then there is no user and thus it returns false
	if result == (models.User{}) {
		return false
	}

	//if result is not empty then there is a user present with the same email id and thus returns true
	return true
}

//probably in a different file
