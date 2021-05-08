package redis

type Redis interface {
	HGet(k string, v string) string
	HSet(k string, kv ...string)
}

type Client interface {
	Save(k, v string)
	Load(k string) string
}

func New(client Redis) Client {
	return &impl{client: client}
}

type impl struct {
	client Redis
}

const hash = "this_is_sparta"

func (i *impl) Save(k, v string) {
	i.client.HSet(hash, k, v)
}

func (i *impl) Load(k string) string {
	return i.client.HGet(hash, k)
}
