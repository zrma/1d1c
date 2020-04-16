package tutorial30daysofcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	t.Run("https://www.hackerrank.com/challenges/30-queues-stacks/problem", func(t *testing.T) {
		t.Run("queue 구조체는 큐 자료구조의 기능을 충실히 수행한다.", func(t *testing.T) {
			q := queue{}

			q.enqueue("a")
			q.enqueue("b")
			q.enqueue("c")

			actual, ok := q.dequeue()
			assert.True(t, ok)
			assert.Equal(t, "a", actual)

			actual, ok = q.dequeue()
			assert.True(t, ok)
			assert.Equal(t, "b", actual)

			actual, ok = q.dequeue()
			assert.True(t, ok)
			assert.Equal(t, "c", actual)

			_, ok = q.dequeue()
			assert.False(t, ok)
		})

		t.Run("stack 구조체는 스택 자료구조의 기능을 충실히 수행한다.", func(t *testing.T) {
			s := stack{}

			s.push("a")
			s.push("b")
			s.push("c")

			actual, ok := s.pop()
			assert.True(t, ok)
			assert.Equal(t, "c", actual)

			actual, ok = s.pop()
			assert.True(t, ok)
			assert.Equal(t, "b", actual)

			actual, ok = s.pop()
			assert.True(t, ok)
			assert.Equal(t, "a", actual)

			_, ok = s.pop()
			assert.False(t, ok)
		})
	})
}
