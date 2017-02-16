package model

import (
	"gopkg.in/redis.v5"
)

func NewClient() *redis.Client{
	client := redis.NewClient(&redis.Options{
		Addr:     "address",
		Password: "password",
		DB:       0,
	})
	return client
}
