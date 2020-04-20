package dp_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Dp Suite")
}
