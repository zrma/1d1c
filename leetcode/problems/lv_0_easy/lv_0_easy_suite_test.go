package lv_0_easy_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLv0Easy(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Lv0Easy Suite")
}
