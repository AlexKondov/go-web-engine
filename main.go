package main

import (
	"log"
	"net/http"

	"basaigbook/controllers"
)

func main() {
	api := &controllers.API{}

	if err := http.ListenAndServe(":8080", api); err != nil {
		log.Println(err)
	}
}
