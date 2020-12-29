package main

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatal(err) // handle error
	}
	defer conn.Close()

	_, err = conn.Do(
		"HMSET",
		"book:1",
		"fabio", "sierra",
	)

	title, err := redis.String(conn.Do("HGET", "book:1", "fabio"))
	if err != nil {
		log.Fatal(err) // handle error
	}
	fmt.Println("Book Title:", title)
}
