package controller

import (
	"GoTools/model"
	"context"
	"encoding/json"
	"log"

	"github.com/go-redis/redis/v8"
)

func GetUsers() []model.User {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	var ctx = context.Background()

	value, err := client.Get(ctx, "users").Result()
	if err != nil {
		log.Println("Redis Error")
		log.Println(err)
		return nil
	}

	var users []model.User
	_ = json.Unmarshal([]byte(value), &users)

	return users
}

func SetUsers(users []model.User) {
	converted, err := json.Marshal(users)
	if err != nil {
		log.Println(err)
		return
	}

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	var ctx = context.Background()

	err = client.Set(ctx, "users", converted, 0).Err()
	if err != nil {
		log.Println("Redis Failed")
		log.Println(err)
		return
	} else {
		log.Println("Redis Success")
	}
}
