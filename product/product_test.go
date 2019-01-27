package product_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/terminatingcode/gilded-rose/item"
	. "github.com/terminatingcode/gilded-rose/product"
)

var _ = Describe("Product", func() {
	Describe("DefaultVariant", func() {
		It("degrades in quality by 1", func() {
			def := Product{
				Item:    item.New("", 1, 1),
				Variant: DefaultVariant{},
			}
			def.Update()
			Expect(def.Item.Quality()).To(Equal(0))
			Expect(def.Item.SellIn()).To(Equal(0))
		})
	})

	Describe("ConjuredVariant", func() {
		It("degrades in quality by 2", func() {
			conjured := Product{
				Item:    item.New("", 1, 4),
				Variant: ConjuredVariant{},
			}
			conjured.Update()
			Expect(conjured.Item.SellIn()).To(Equal(0))
			Expect(conjured.Item.Quality()).To(Equal(2))
		})
	})

	Describe("AgedBrieVariant", func() {
		It("increases in quality by 1", func() {
			brie := Product{
				Item:    item.New("", 1, 4),
				Variant: AgedBrieVariant{},
			}
			brie.Update()
			Expect(brie.Item.SellIn()).To(Equal(0))
			Expect(brie.Item.Quality()).To(Equal(5))
		})
	})

	Describe("BackStagePassVariant", func() {
		It("increases in quality by 1 before sellIn is 10 days away", func() {
			pass := Product{
				Item:    item.New("", 15, 4),
				Variant: BackStagePassVariant{},
			}
			pass.Update()
			Expect(pass.Item.SellIn()).To(Equal(14))
			Expect(pass.Item.Quality()).To(Equal(5))
		})

		It("increases in quality by 2 if sellIn is less than 10 days away", func() {
			pass := Product{
				Item:    item.New("", 10, 4),
				Variant: BackStagePassVariant{},
			}
			pass.Update()
			Expect(pass.Item.SellIn()).To(Equal(9))
			Expect(pass.Item.Quality()).To(Equal(6))
		})

		It("increases in quality by 3 if sellIn is less than 5 days away", func() {
			pass := Product{
				Item:    item.New("", 5, 4),
				Variant: BackStagePassVariant{},
			}
			pass.Update()
			Expect(pass.Item.SellIn()).To(Equal(4))
			Expect(pass.Item.Quality()).To(Equal(7))
		})
	})

	Describe("SulfurasVariant", func() {
		It("will never decrease in sellIn nor quality", func() {
			sulf := Product{
				Item:    item.New("", 15, 4),
				Variant: SulfurasVariant{},
			}
			sulf.Update()
			Expect(sulf.Item.SellIn()).To(Equal(15))
			Expect(sulf.Item.Quality()).To(Equal(4))
		})
	})
})
