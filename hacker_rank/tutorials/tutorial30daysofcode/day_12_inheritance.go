package tutorial30daysofcode

import "fmt"

type persona struct {
	firstName string
	lastName  string
	id        int
}

func newPerson(firstName, lastName string, identification int) *persona {
	return &persona{firstName, lastName, identification}
}

func (p persona) printPerson() string {
	return fmt.Sprintf("Name: %s, %s\nID: %d", p.lastName, p.firstName, p.id)
}

type student struct {
	*persona

	testScores []int
}

func newStudent(firstName, lastName string, identification int, scores []int) *student {
	return &student{
		newPerson(firstName, lastName, identification),
		scores,
	}
}

func (s student) calculate() string {
	length := len(s.testScores)
	sum := 0
	for _, s := range s.testScores {
		sum += s
	}

	avr := sum / length
	switch {
	case avr >= 90:
		return "O"
	case avr >= 80:
		return "E"
	case avr >= 70:
		return "A"
	case avr >= 55:
		return "P"
	case avr >= 40:
		return "D"
	default:
		return "T"
	}
}
