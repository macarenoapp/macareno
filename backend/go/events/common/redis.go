package common

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"context"
	
)
//"fmt"
var ctx = context.Background()

//S_Users :  Struct of the redis package
type S_Users struct{
	Key   string `json:"key"   binding:"required"`
	Value string `json:"value" binding:"required"`
}

//NewClient :  Creates a new client to use Redis
func NewClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	return rdb 
}

func SaveUser(email string, value string, client *redis.Client) error {
	
	err := client.Set(ctx,"email", email, 0).Err()

	if err != nil {
		return err
	}
	return nil
}


func GetUser(key string, client *redis.Client) string {

	email, err := client.Get(ctx, key).Result() 

	if err != nil {
		fmt.Println("ERROR en GetUser de redis: ", err)
	}
	return email
}