package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/vishnumenon/budgetapplication/database"
	"github.com/vishnumenon/budgetapplication/routers"
)

func main() {
	fmt.Println("Budget Application API")
	r := routers.Router()

	port := database.GetEnvValue("PORT")
	fmt.Println("Server Starting....", port)
	log.Fatal(http.ListenAndServe(port, r))
	fmt.Println("Server Listening at PORT: 4000")

}
