package main

import (
	"fmt"
	"golang-redis/routers"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Start REDIS API")
	routers.HandleRouters()
	log.Fatal(http.ListenAndServe(":8080", routers.R))
}
