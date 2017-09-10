package lock

import "github.com/mediocregopher/radix.v2/redis"

const lock_key = "LOCK_KEY"
const lock_value = "LOCK_VALUE"

func acquire(client *redis.Client) bool {
	current_value, _ := client.Cmd("GET", lock_key).Str()
	if current_value == "" {
		client.Cmd("SET", lock_key, lock_value)
		return true
	} else {
		return false
	}
}

func release(client *redis.Client) {
	client.Cmd("DEL", lock_key)
}
