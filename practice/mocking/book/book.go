package book

//go:generate go run github.com/golang/mock/mockgen -destination=../mocks/mock_book.go -package=mocks . Book
type Book interface {
	Page(string) int
	Read(string) bool
}

func PlayWithBook(b Book, name string) (int, bool) {
	return b.Page(name), b.Read(name)
}
