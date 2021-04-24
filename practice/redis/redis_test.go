package redis

import (
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/stretchr/testify/assert"
)

func TestRedisHGetHSet(t *testing.T) {
	mock, err := miniredis.Run()
	assert.NoError(t, err)
	assert.NotNil(t, mock)

	t.Cleanup(func() {
		mock.Close()
	})

	const (
		hash = "hash123"
		key  = "key1"
		want = "val2"
	)

	client := NewClient(mock)
	client.HSet(hash, key, want)

	got := client.HGet(hash, key)
	assert.Equal(t, want, got)
}
