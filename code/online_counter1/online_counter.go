package online_counter

import "github.com/mediocregopher/radix.v2/redis"

const online_user_set = "ONLINE_USER_SET"

func set_online(client *redis.Client, user string) {
	client.Cmd("SADD", online_user_set, user)
}

func count_online(client *redis.Client) int64 {
	repl, _ := client.Cmd("SCARD", online_user_set).Int64()
	return repl
}

func is_online_or_not(client *redis.Client, user string) bool {
	repl, _ := client.Cmd("SISMEMBER", online_user_set, user).Int()
	return repl == 1
}
