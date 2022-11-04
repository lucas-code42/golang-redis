package controller

import (
	"encoding/json"
	"fmt"
	"golang-redis/model"
	"golang-redis/redis"
	"io"
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
	var sessionKey model.SessionModel
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	if err = json.Unmarshal(body, &sessionKey); err != nil {
		log.Fatal(err)
	}

	if redis.HowManyRequests(sessionKey.Key) >= 10 {
		w.Write([]byte("Your requests has been blocked"))
		return
	}

	countRequests, err := redis.IncrementRequests(sessionKey.Key)
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte(countRequests))
}

// TimeLeft mostrará quanto tempo de sessao ainda resta
func TimeLeft(w http.ResponseWriter, r *http.Request) {
	var sessionKey model.SessionModel
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	if err = json.Unmarshal(body, &sessionKey); err != nil {
		log.Fatal(err)
	}

	ttl := redis.GetTTL(sessionKey.Key)

	w.Write([]byte(fmt.Sprintf("Time left session %s", ttl)))

}
