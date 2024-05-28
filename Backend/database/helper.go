package database

import (
	"context"
	"fmt"
	"log"

	"github.com/vishnumenon/budgetapplication/models"
)

func AddUser(user models.User) {
	inserted, err := UserCollection.InsertOne(context.Background(), user)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("User has been inserted Successfully with ID: ", inserted.InsertedID)
}
