package item_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/terminatingcode/gilded-rose/item"
)

var _ = Describe("Item", func() {
	Describe("UpdateQuality", func() {
		It("will degrade twice as fast when sellIn less than 0", func() {
			item := New("", 0, 4)
			item.UpdateQuality(-1)
			Expect(item.Quality()).To(Equal(2))
		})

		It("will not degrade past 0", func() {
			item := New("", 0, 0)
			item.UpdateQuality(-1)
			Expect(item.Quality()).To(Equal(0))
		})

		It("will not increase past 50", func() {
			item := New("", 0, 49)
			item.UpdateQuality(2)
			Expect(item.Quality()).To(Equal(50))
		})
	})

	Describe("given a sulfuras", func() {
		It("will not degrade in quality or decrement sellIn", func() {
			item := New("Sulfuras, Hand of Ragnaros", 1, 1)
			err := UpdateItems([]*Item{item})
			Expect(err).NotTo(HaveOccurred())
			Expect(item.SellIn()).To(Equal(1))
			Expect(item.Quality()).To(Equal(1))

		})
	})

	Describe("given a backstage pass", func() {
		It("increases in quality by 2 if sellIn <=10", func() {
			item := New("Backstage passes to a TAFKAL80ETC concert", 10, 1)
			err := UpdateItems([]*Item{item})
			Expect(err).NotTo(HaveOccurred())
			Expect(item.Quality()).To(Equal(3))
		})

		It("increases in quality by 3 if sellIn <=5", func() {
			item := New("Backstage passes to a TAFKAL80ETC concert", 5, 1)
			err := UpdateItems([]*Item{item})
			Expect(err).NotTo(HaveOccurred())
			Expect(item.Quality()).To(Equal(4))
		})

		It("loses all quality if sellIn <=0", func() {
			item := New("Backstage passes to a TAFKAL80ETC concert", 0, 1)
			err := UpdateItems([]*Item{item})
			Expect(err).NotTo(HaveOccurred())
			Expect(item.Quality()).To(Equal(0))
		})
	})

	Describe("given a conjured item", func() {
		It("degrades in quality by 2", func() {
			item := New("Conjured item", 1, 2)
			err := UpdateItems([]*Item{item})
			Expect(err).NotTo(HaveOccurred())
			Expect(item.Quality()).To(Equal(0))
		})

		It("degrades to 0 even if odd number", func() {
			item := New("Conjured item", 1, 1)
			err := UpdateItems([]*Item{item})
			Expect(err).NotTo(HaveOccurred())
			Expect(item.Quality()).To(Equal(0))
		})
	})
})
