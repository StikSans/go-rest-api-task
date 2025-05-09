package main

import (
	"log"
	"net/http"
	"task-manager/internal/delivery"
)

func main() {
	router := delivery.NewRouter()

  log.Println("Server is running on port 8080")

  if err := http.ListenAndServe(":8080", router); err != nil {
    log.Fatal(err)
  }

}