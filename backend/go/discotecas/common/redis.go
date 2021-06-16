package common

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"context"
	"strconv"
	// "reflect"
	
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

func SaveUserLike(user uint, discoteca string, client *redis.Client) error{

	var newUsers string;

	//Pillamos el total de likes que tiene esa discoteca
	totalUsers, err2 := client.Get(ctx, discoteca).Result()

	if err2 != nil{
		fmt.Println("ERR2: ", err2)
	}

	newUsers = totalUsers +","+ strconv.Itoa(int(user)) //Le paso un uint, pero quiere un int

	err3 := client.Set(ctx,discoteca, newUsers, 0).Err()

	if err3 != nil{
		fmt.Println("ERRRORRRR: ",err3)
		return err3
	}

	return nil

}

func DeleteDiscoRedis(id string, client *redis.Client) error{

	err := client.Del(ctx, id).Err()

	if err != nil{
		fmt.Println("Error deletind disco from redis")
		return err
	}

	fmt.Println("Delete from REDIS OK")

	return nil

}