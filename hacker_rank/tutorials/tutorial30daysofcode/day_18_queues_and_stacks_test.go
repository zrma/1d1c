package tutorial30daysofcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-queues-stacks/problem")

	q := queue{}

	q.enqueue("a")
	q.enqueue("b")
	q.enqueue("c")

	for i, tt := range []struct {
		ok   bool
		want string
	}{
		{true, "a"},
		{true, "b"},
		{true, "c"},
		{false, ""},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got, ok := q.dequeue()
			assert.Equal(t, tt.ok, ok)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestStack(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-queues-stacks/problem")

	s := stack{}

	s.push("a")
	s.push("b")
	s.push("c")

	for i, tt := range []struct {
		ok   bool
		want string
	}{
		{true, "c"},
		{true, "b"},
		{true, "a"},
		{false, ""},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got, ok := s.pop()
			assert.Equal(t, tt.ok, ok)
			assert.Equal(t, tt.want, got)
		})
	}
}
