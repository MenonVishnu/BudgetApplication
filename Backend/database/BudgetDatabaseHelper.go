package database

import (
	"context"
	"fmt"
	"log"

	"github.com/vishnumenon/budgetapplication/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddBudget(budget models.Budget) primitive.ObjectID {
	inserted, err := BudgetCollection.InsertOne(context.Background(), budget)

	if err != nil {
		log.Fatal(err)
	}
	//changing data type from interface to primitive.objectId
	budgetId := inserted.InsertedID.(primitive.ObjectID)

	fmt.Println("User has been inserted Successfully with ID: ", budgetId)
	return budgetId
}
