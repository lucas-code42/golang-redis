package controller

import (
	"encoding/json"
	"fmt"
	"golang-redis/redis"
	"log"
	"net/http"
)

// Session vai mostrar um ID de sessao
func Session(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	
	userCredentials, err := redis.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	
	userCredentialsJSON, err := json.Marshal(userCredentials)
	if err != nil {
		log.Fatal(err)
	}
	
	w.Write(userCredentialsJSON)
}

// Requests retornará quantas requisições o usuario da sessão ja fez
func Requests(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Requests")
}

// TimeLeft mostrará quanto tempo de sessao ainda resta
func TimeLeft(w http.ResponseWriter, r *http.Request) {
	fmt.Println("TimeLeft")
}
