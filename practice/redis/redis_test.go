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
		key  = "key1"
		want = "val2"
	)

	client := New(mock)
	client.Save(key, want)

	got := client.Load(key)
	assert.Equal(t, want, got)
}
