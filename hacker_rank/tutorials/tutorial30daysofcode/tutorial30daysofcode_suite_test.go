package tutorial30daysofcode_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTutorial30daysofcode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tutorial30daysofcode Suite")
}
