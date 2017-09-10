package main

import (
    "fmt"
	"github.com/mediocregopher/radix.v2/redis"
)

const online_user_hll = "ONLINE_USER_HLL"

func set_online(client *redis.Client, user string) {
    client.Cmd("PFADD", online_user_hll, user)
}

func count_online(client *redis.Client) int64 {
    repl, _ := client.Cmd("PFCOUNT", online_user_hll).Int64()
    return repl
}
