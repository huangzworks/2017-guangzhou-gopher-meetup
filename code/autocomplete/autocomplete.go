package autocomplete

import "github.com/mediocregopher/radix.v2/redis"

const autocomplete = "autocomplete::"

func feed(client *redis.Client, content string, weight int) {
	for i, _ := range content {
		segment := content[:i+1]
		key := autocomplete + segment
		client.Cmd("ZINCRBY", key, weight, content)
	}
}

func hint(client *redis.Client, prefix string, count int) []string {
	key := autocomplete + prefix
	result, _ := client.Cmd("ZREVRANGE", key, 0, count-1).List()
	return result
}
