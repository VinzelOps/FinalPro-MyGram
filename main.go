package main

import (
	"github.com/FinalPro-MyGram/database"
	"github.com/FinalPro-MyGram/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	err := r.Run()
	if err != nil {
		return
	}
}
