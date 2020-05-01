package lv2hard_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLv2Hard(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Lv2Hard Suite")
}
