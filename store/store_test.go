package store_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/terminatingcode/gilded-rose/item"
	p "github.com/terminatingcode/gilded-rose/product"
	. "github.com/terminatingcode/gilded-rose/store"
)

var _ = Describe("Store", func() {
	It("will update the inventory", func() {
		defProduct := p.Product{item.New("", 1, 1), p.DefaultVariant{}}
		conjProduct := p.Product{item.New("", 1, 2), p.ConjuredVariant{}}
		brieProduct := p.Product{item.New("", 1, 1), p.AgedBrieVariant{}}
		backProduct := p.Product{item.New("", 1, 1), p.BackStagePassVariant{}}
		sulfProduct := p.Product{item.New("", 1, 1), p.SulfurasVariant{}}
		store := Store{
			Products: []p.Product{defProduct, conjProduct, brieProduct, backProduct, sulfProduct},
		}
		store.UpdateInventory(1)
		Expect(store.Products[0].Item.SellIn()).To(Equal(0))
		Expect(store.Products[0].Item.Quality()).To(Equal(0))
		Expect(store.Products[1].Item.SellIn()).To(Equal(0))
		Expect(store.Products[1].Item.Quality()).To(Equal(0))
		Expect(store.Products[2].Item.SellIn()).To(Equal(0))
		Expect(store.Products[2].Item.Quality()).To(Equal(2))
		Expect(store.Products[3].Item.SellIn()).To(Equal(0))
		Expect(store.Products[3].Item.Quality()).To(Equal(4))
		Expect(store.Products[4].Item.SellIn()).To(Equal(1))
		Expect(store.Products[4].Item.Quality()).To(Equal(1))

	})
})
