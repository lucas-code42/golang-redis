package redis

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/google/uuid"
)

func redisConnection() (*redis.Client, error) {
	rd := redis.NewClient(&redis.Options{})
	pong := rd.Ping()

	if pong.Val() == "PONG" {
		return rd, nil
	} else {
		return nil, fmt.Errorf("error to connect to redis")
	}
}

// CreateSession cria sessao de um usuario
func CreateSession() (map[string]string, error) {
	rd, err := redisConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer rd.Close()

	sessionID := uuid.New()
	userID := uuid.New()
	thirtyMinutes := time.Duration(time.Minute * 30)

	key := sessionID.String()
	value := userID.String()

	create := rd.HSet(key, "userID", value)
	expire := rd.Expire(key, thirtyMinutes)

	createStatus, err := create.Result()
	if err != nil {
		log.Fatal(err)
	}
	expireStatus, err := expire.Result()
	if err != nil {
		log.Fatal(err)
	}

	var returnStmt map[string]string
	if createStatus && expireStatus {
		fmt.Println("deu certo")
		returnStmt = map[string]string{"session": key, "userID": value}
		return returnStmt, nil
	} else {
		returnStmt = map[string]string{"session": "", "userID": ""}
		return nil, fmt.Errorf("error to create session")
	}
}
