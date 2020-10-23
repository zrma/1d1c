package book_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	. "github.com/zrma/1d1go/practice/mocking/book"
	"github.com/zrma/1d1go/practice/mocking/mocks"
)

func TestPlayWithBook(t *testing.T) {
	t.Log("gomock 사용 연습")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockBook(ctrl)

	m.EXPECT().Read(gomock.Eq("a")).Return(true).AnyTimes()
	m.EXPECT().Page(gomock.Eq("a")).Return(123).AnyTimes()

	m.EXPECT().Read(gomock.Eq("b")).Return(false).AnyTimes()
	m.EXPECT().Page(gomock.Eq("b")).Return(0).AnyTimes()

	t.Run("name a", func(t *testing.T) {
		cnt, ok := PlayWithBook(m, "a")
		assert.True(t, ok)
		assert.Equal(t, cnt, 123)
	})

	t.Run("name b", func(t *testing.T) {
		cnt, ok := PlayWithBook(m, "b")
		assert.False(t, ok)
		assert.Equal(t, cnt, 0)
	})
}
