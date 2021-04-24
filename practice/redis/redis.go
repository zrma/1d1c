package redis

type Redis interface {
	HGet(k string, v string) string
	HSet(k string, kv ...string)
}

func NewClient(client Redis) Redis {
	return &redisClient{client: client}
}

type redisClient struct {
	client Redis
}

func (r redisClient) HGet(hash string, key string) string {
	return r.client.HGet(hash, key)
}

func (r redisClient) HSet(hash string, kv ...string) {
	r.client.HSet(hash, kv...)
}
