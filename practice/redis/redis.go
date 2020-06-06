package redis

type Redis interface {
	HGet(k string, f string) string
	HSet(k string, fv ...string)
}

func HGet(redis Redis, hash, key string) string {
	return redis.HGet(hash, key)
}

func HSet(redis Redis, hash string, fv ...string) {
	redis.HSet(hash, fv...)
}
