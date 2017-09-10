package lock

import "github.com/mediocregopher/radix.v2/redis"

const lock_key = "LOCK_KEY"
const lock_value = "LOCK_VALUE"

func acquire(client *redis.Client) bool {
	client.Cmd("WATCH", lock_key)
	current_value, _ := client.Cmd("GET", lock_key).Str()
	if current_value == "" {
		client.Cmd("MULTI")
		client.Cmd("SET", lock_key, lock_value)
		repl, _ := client.Cmd("EXEC").List()
		if repl != nil {
			return true
		}
	}
	return false
}

func release(client *redis.Client) {
	client.Cmd("DEL", lock_key)
}
