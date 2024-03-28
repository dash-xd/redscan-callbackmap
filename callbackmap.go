package callbacks

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var CallbackMap = map[string]CallbackFunc{
	"SaveSubscriptionGroup": func(client *redis.Client, key string) ([]byte, error) {
		asubId, channelName := parseKey(key)
		clientId, err := client.Do(context.Background(), "Client", "Id").Result()
		if err != nil {
			return nil, err
		}
		return Callbacks().callLuaFunction(client, "FCALL", "Registrar", 0, "subg", clientId, asubId, channelName)(key)
	},

	"Delete": func(client *redis.Client, key string) ([]byte, error) {
		res, err := client.Del(context.Background(), key).Result()
		var resBytes []byte
		if err != nil {
			resBytes = []byte(fmt.Sprintf("Failed to delete key: %s", err.Error()))
		} else {
			resBytes = []byte(fmt.Sprintf("Number of keys deleted: %d", res))
		}
		return resBytes, err
	},

	"Mock": func(client *redis.Client, key string) ([]byte, error) {
		fmt.Println(key)
		return []byte{}, nil
	},
}
