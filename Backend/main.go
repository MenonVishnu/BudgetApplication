package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/vishnumenon/budgetapplication/routers"
)


func main() {
	fmt.Println("Budget Application API")
	r := routers.Router()
	fmt.Println("Server Starting....")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Server Listening at PORT: 4000")

}



