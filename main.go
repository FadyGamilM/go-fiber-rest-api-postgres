package main

import (
	"fmt"
	"net/http"

	"github.com/FadyGamilM/gomongo/database"
	"github.com/FadyGamilM/gomongo/routers"
)

func main() {
	r := routers.Router()
	// connect to the database
	database.Init()

	http.ListenAndServe("localhost:4000", r)
	fmt.Println("Server is up and running on port 4000")
}
