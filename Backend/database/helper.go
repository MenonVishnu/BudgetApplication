package database

import (
	"context"
	"fmt"
	"log"

	"github.com/vishnumenon/budgetapplication/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

func AddUser(user models.User) primitive.ObjectID {

	inserted, err := UserCollection.InsertOne(context.Background(), user)

	if err != nil {
		log.Fatal(err)
	}
	//changing data type from interface to primitive.objectId
	userId := inserted.InsertedID.(primitive.ObjectID)

	fmt.Println("User has been inserted Successfully with ID: ", userId)
	return userId
}

func UpdateUser(user models.User, userId string) {
	id, err := primitive.ObjectIDFromHex(userId)

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

	fmt.Println("User Updated with object ID: ", userId, " Documents Affected: ", updatedUser.ModifiedCount)
}

func DeleteUser(userId string){
	id, err := primitive.ObjectIDFromHex(userId)

	if err!=nil{
		log.Fatal(err)
	}

	filter := bson.M{"_id":id}

	deletedUser, err := UserCollection.DeleteOne(context.Background(),filter)

	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println("User Deleted with ObjectId: ",userId, " Documents Affected: ", deletedUser.DeletedCount)

}

func DeleteAllUser(){
	filter := bson.M{}
	deletedUser, err := UserCollection.DeleteMany(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("All Users Deleted, Documents Affected:  ",deletedUser.DeletedCount)

}

//possibly will not work
func GetUser(userId string) models.User{

	var user models.User

	id, err := primitive.ObjectIDFromHex(userId)
	if err!=nil{
		log.Fatal(err)
	}

	filter := bson.M{"_id":id}
	project := bson.M{"password": 0}
	opts := options.FindOne().SetProjection(project)

	err = UserCollection.FindOne(context.Background(),filter, opts).Decode(&user)

	//please check will this work or not??
	if err == nil {
		return models.User{}
	}

	return user
}

func GetAllUser() []primitive.M{
	var users []primitive.M 

	filter := bson.M{}
	project := bson.M{"password": 0}
	opts := options.Find().SetProjection(project)

	curr, err = UserCollection.Find(context.Background(),filter, opts)

	
}

// helper function
// the below function returns true if there exists a user with the email-id provided in the function parameters
func CheckUser(email string) bool {

	var result models.User

	filter := bson.M{"email": email}
	project := bson.M{"password": 0}
	opts := options.FindOne().SetProjection(project)

	err := UserCollection.FindOne(context.Background(), filter, opts).Decode(&result)

	if err != nil {
		// log.Fatal(err)
		return false
	}

	//if result is empty then there is no user and thus it returns false
	if result == (models.User{}) {
		return false
	}

	//if result is not empty then there is a user present with the same email id and thus returns true
	return true
}

// Encrypting Password
func EncryptPassword(password string) string {
	passwordByte := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordByte, bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hashedPassword)
}

// Comparing Given password with already present password
func CheckPassword(hashedPassword string, givenPassword string) bool {
	hashedPasswordByte := []byte(hashedPassword)
	givenPasswordByte := []byte(givenPassword)

	err := bcrypt.CompareHashAndPassword(hashedPasswordByte, givenPasswordByte)
	if err != nil {
		return false
	}
	return true
}
