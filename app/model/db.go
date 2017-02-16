package model

import (
	"gopkg.in/redis.v5"
)

func NewClient() *redis.Client{
	client := redis.NewClient(&redis.Options{
		Addr:     "porgy.redistogo.com:10051",
		Password: "7e7dc2e978d297c6851fcd8344f46b39",
		DB:       0,  // DB
	})
	return client
}
