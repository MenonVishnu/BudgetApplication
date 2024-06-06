package database

import (
	"context"
	"fmt"
	"log"

	"github.com/vishnumenon/budgetapplication/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// add budget to database
func AddBudget(budget models.Budget) primitive.ObjectID {
	inserted, err := BudgetCollection.InsertOne(context.Background(), budget)

	if err != nil {
		log.Fatal(err)
	}
	//changing data type from interface to primitive.objectId
	budgetId := inserted.InsertedID.(primitive.ObjectID)

	fmt.Println("Budget has been inserted Successfully with ID: ", budgetId)
	return budgetId
}

// update budget to database
func UpdateBudget(budget models.Budget, budgetId string) {
	id, err := primitive.ObjectIDFromHex(budgetId)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{
		"amount": budget.Amount,
		"tags":   budget.Tags,
		"Date":   budget.Date,
		"User":   budget.User,
	}}

	updatedBudget, err := BudgetCollection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Budget Updated with object ID: ", budgetId, " Documents Affected: ", updatedBudget.ModifiedCount)

}

// admin route - delete all budgets
func DeleteAllBudget() {
	filter := bson.M{}

	deleted, err := BudgetCollection.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("All Budgets Deleted by Admin. Documents Affected: ", deleted.DeletedCount)
}

// user route
func DeleteBudget(budgetId string) {
	id, err := primitive.ObjectIDFromHex(budgetId)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}

	deleted, err := BudgetCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Budget Deleted with object ID: ", budgetId, " Documents Affected: ", deleted.DeletedCount)
}


// delete all budget of a specific user -> for admin and user
// not completed / may not work
func DeleteAllUsersBudget(givenUserId string) {
	userId, err := primitive.ObjectIDFromHex(givenUserId)
	if err != nil {
		log.Fatal(err)
	}

	// models.Budget
	// var budgets []models.Budget
	//filter out based on user id
	filter := bson.M{"user._id": userId}
	deletedBudgets, err := BudgetCollection.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Budget belonging to User ID: ", userId, " Deleted Successfully. ", "Documents Affected: ", deletedBudgets.DeletedCount)
}

// admin get all budgets of all users
func GetAllBudgets() []models.Budget{

	var budgets []models.Budget

	filter := bson.M{}
	//object Id will not be given
	// opts := options.Find().SetProjection(bson.M{"_id":0})

	//TODO: filter out password from database using projection
	curr, err := BudgetCollection.Find(context.Background(), filter)

	if err!=nil{
		log.Fatal(err)
	}

	for curr.Next(context.Background()){
		var budget models.Budget
		err := curr.Decode(&budget)
		if err!=nil{
			log.Fatal(err)
		}
		budgets = append(budgets, budget)
	}

	return budgets
}

//user: get one budget based on the id
func GetOneBudget(budgetId string) models.Budget{
	id, err := primitive.ObjectIDFromHex(budgetId)

	if err!=nil{
		log.Fatal(err)
	}

	var budget models.Budget

	filter := bson.M{"_id":id}
	//TODO: filter out password from database using projection
	err = 	BudgetCollection.FindOne(context.Background(), filter).Decode(&budget)
	if err!=nil{
		return models.Budget{}
	}
	return budget
}

//user: get all budgets of a specific user using userId 
func GetAllUserBudget(userId string) []models.Budget{
	id, err := primitive.ObjectIDFromHex(userId)
	if err!=nil{
		log.Fatal(err)
	}

	var budgets []models.Budget

	filter := bson.M{"user._id":id}
	//TODO: filter out password from database using projection
	
	curr, err := BudgetCollection.Find(context.Background(), filter)
	if err!=nil{
		log.Fatal(err)
	}

	
	for curr.Next(context.Background()){
		var budget models.Budget
		err = curr.Decode(&budget)
		if err!=nil{
			log.Fatal(err)
		}
		budgets = append(budgets, budget)
	}

	return budgets
}