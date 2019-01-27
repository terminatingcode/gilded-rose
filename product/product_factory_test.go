package product_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/terminatingcode/gilded-rose/product"
)

var _ = Describe("ProductFactory", func() {
	It("creates a default product", func() {
		p := NewProductFactory("", 1, 1)
		Expect(p.Variant).To(Equal(DefaultVariant{}))
	})

	It("creates a conjured product", func() {
		p := NewProductFactory("Conjured ", 1, 1)
		Expect(p.Variant).To(Equal(ConjuredVariant{}))
	})

	It("creates an aged brie product", func() {
		p := NewProductFactory("Aged Brie", 1, 1)
		Expect(p.Variant).To(Equal(AgedBrieVariant{}))
	})

	It("creates a sulfuras product", func() {
		p := NewProductFactory("Sulfuras hand of whoever", 1, 1)
		Expect(p.Variant).To(Equal(SulfurasVariant{}))
	})

	It("creates a backstage pass product", func() {
		p := NewProductFactory("Backstage passes to your office party, i.e. you have been volunteered for cleanup!", 1, 1)
		Expect(p.Variant).To(Equal(BackStagePassVariant{}))
	})
})
