package main

import (
	"fmt"
	"github.com/mediocregopher/radix.v2/redis"
)

func main() {
	client, _ := redis.Dial("tcp", "localhost:6379")
	defer client.Close()

	repl := client.Cmd("PING")
	result, _ := repl.Str()

	fmt.Println(result) // "PONG"
}
