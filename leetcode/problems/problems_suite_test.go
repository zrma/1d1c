package problems_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestProblems(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Problems Suite")
}
