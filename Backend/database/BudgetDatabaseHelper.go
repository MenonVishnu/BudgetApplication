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

func DeleteBudget(budgetId string) {
	id, err := primitive.ObjectIDFromHex(budgetId)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}

	deleted, err := BudgetCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal()
	}
	fmt.Println("Budget Deleted with object ID: ", budgetId, " Documents Affected: ", deleted.DeletedCount)
}
