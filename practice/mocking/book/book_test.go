package book_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	. "1d1go/practice/mocking/book"
	"1d1go/practice/mocking/mocks"
)

func TestPlayWithBook(t *testing.T) {
	t.Parallel()

	//noinspection SpellCheckingInspection
	t.Run("gomock 사용 연습", func(t *testing.T) {
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
	})

	t.Run("testify.mock 사용 연습", func(t *testing.T) {
		m := new(myMockBook)
		m.On("Read", "a").Return(true)
		m.On("Page", "a").Return(123)

		m.On("Read", "b").Return(false)
		m.On("Page", "b").Return(0)

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
	})
}

type myMockBook struct {
	mock.Mock
}

func (m *myMockBook) Page(s string) int {
	args := m.Called(s)
	return args.Int(0)
}

func (m *myMockBook) Read(s string) bool {
	args := m.Called(s)
	return args.Bool(0)
}
