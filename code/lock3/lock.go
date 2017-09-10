package lock

import "github.com/mediocregopher/radix.v2/redis"

const lock_key = "LOCK_KEY"
const lock_value = "LOCK_VALUE"

func acquire(client *redis.Client) bool {
	repl, _ := client.Cmd("SET", lock_key, lock_value, "NX").Str()
	return repl != ""
}

func release(client *redis.Client) {
	client.Cmd("DEL", lock_key)
}
