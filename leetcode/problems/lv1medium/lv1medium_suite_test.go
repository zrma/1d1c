package lv1medium_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLv1Medium(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Lv1Medium Suite")
}
