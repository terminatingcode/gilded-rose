package item_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGildedRose(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GildedRose Suite")
}