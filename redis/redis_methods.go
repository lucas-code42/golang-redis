package redis

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/google/uuid"
)

// redisConnection faz a conexao com redis e a retorna
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
	var countRequests int
	countRequests += 1
	requests := rd.HSet(key, "requests", strconv.Itoa(countRequests))
	expire := rd.Expire(key, thirtyMinutes)

	createStatus, err := create.Result()
	if err != nil {
		log.Fatal(err)
	}
	requestsStatus, err := requests.Result()
	if err != nil {
		log.Fatal(err)
	}
	expireStatus, err := expire.Result()
	if err != nil {
		log.Fatal(err)
	}

	var returnStmt map[string]string
	if createStatus && expireStatus && requestsStatus {
		fmt.Println("deu certo")
		returnStmt = map[string]string{"session": key, "userID": value}
		return returnStmt, nil
	} else {
		returnStmt = map[string]string{"session": "", "userID": ""}
		return nil, fmt.Errorf("error to create session")
	}
}

// IncrementRequests adcionará +1 a cada request
func IncrementRequests(sessionKey string) (string, error) {
	rd, err := redisConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer rd.Close()

	requestValueStr := rd.HGet(sessionKey, "requests")
	requestValueInt, err := strconv.ParseInt(requestValueStr.Val(), 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	requestValueInt += 1

	newValue := strconv.Itoa(int(requestValueInt))
	fmt.Printf("newValue: %v\n", newValue)

	operation := rd.HSet(sessionKey, "requests", newValue)
	if err = operation.Err(); err != nil {
		log.Fatal(err)
	}
	return newValue, nil
}

func HowManyRequests(sessionKey string) int {
	rd, err := redisConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer rd.Close()

	requestValueStr := rd.HGet(sessionKey, "requests")
	requestValueInt, err := strconv.ParseInt(requestValueStr.Val(), 10, 32)
	if err != nil {
		log.Fatal(err)
	}

	return int(requestValueInt)
}

func GetTTL(sessionKey string) string {
	rd, err := redisConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer rd.Close()

	timeLeft := rd.TTL(sessionKey)
	timeLeftValue := timeLeft.Val()
	fmt.Println(timeLeftValue)
	return timeLeftValue.String()
}
