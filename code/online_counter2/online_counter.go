package online_counter

import "github.com/mediocregopher/radix.v2/redis"

const online_user_bitmap = "ONLINE_USER_BITMAP"

func set_online(client *redis.Client, user_id int64) {
	client.Cmd("SETBIT", online_user_bitmap, user_id, 1)
}

func count_online(client *redis.Client) int64 {
	repl, _ := client.Cmd("BITCOUNT", online_user_bitmap).Int64()
	return repl
}

func is_online_or_not(client *redis.Client, user_id int64) bool {
	repl, _ := client.Cmd("GETBIT", online_user_bitmap, user_id).Int()
	return repl == 1
}
