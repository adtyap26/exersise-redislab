package main

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr: "redis-10974.re-cluster1.ps-redislabs.org:10974",
	})
	defer rdb.Close()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to source-db")

	for i := 1; i <= 100; i++ {
		key := fmt.Sprintf("%d", i)
		err = rdb.Set(ctx, key, i, 0).Err()
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Inserted 100 keys into source-db")
}
