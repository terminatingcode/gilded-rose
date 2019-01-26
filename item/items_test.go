package item

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UpdateQuality", func() {
	Describe("given no items", func() {
		It("does not error", func() {
			err := UpdateQuality(nil)
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Describe("given a non special item", func() {
		Context("quality", func() {
			It("will degrade", func() {
				item := Item{
					quality: 1,
				}
				err := UpdateQuality([]*Item{&item})
				Expect(err).NotTo(HaveOccurred())
				Expect(item.quality).To(Equal(0))
			})

			It("will not degrade past 0", func() {
				item := Item{
					quality: 0,
				}
				err := UpdateQuality([]*Item{&item})
				Expect(err).NotTo(HaveOccurred())
				Expect(item.quality).To(Equal(0))
			})

			It("will decrement twice as fast if sellIn less than 0", func() {
				item := Item{
					sellIn:  0,
					quality: 2,
				}
				err := UpdateQuality([]*Item{&item})
				Expect(err).NotTo(HaveOccurred())
				Expect(item.sellIn).To(Equal(-1))
				Expect(item.quality).To(Equal(0))
			})
		})

		Context("sellIn", func() {
			It("will decrement", func() {
				item := Item{
					sellIn: 1,
				}
				err := UpdateQuality([]*Item{&item})
				Expect(err).NotTo(HaveOccurred())
				Expect(item.sellIn).To(Equal(0))
			})
		})
	})

	Describe("Aged Brie", func() {
		It("increases in quality over time", func() {
			item := Item{
				name:    "Aged Brie",
				sellIn:  1,
				quality: 1,
			}
			err := UpdateQuality([]*Item{&item})
			Expect(err).NotTo(HaveOccurred())
			Expect(item.sellIn).To(Equal(0))
			Expect(item.quality).To(Equal(2))
		})

		It("will not increase past 50", func() {
			item := Item{
				name:    "Aged Brie",
				sellIn:  1,
				quality: 50,
			}
			err := UpdateQuality([]*Item{&item})
			Expect(err).NotTo(HaveOccurred())
			Expect(item.sellIn).To(Equal(0))
			Expect(item.quality).To(Equal(50))
		})
	})

	Describe("given a sulfuras", func() {
		It("will not degrade in quality or decrement sellIn", func() {
			item := Item{
				name:    "Sulfuras, Hand of Ragnaros",
				sellIn:  1,
				quality: 1,
			}
			err := UpdateQuality([]*Item{&item})
			Expect(err).NotTo(HaveOccurred())
			Expect(item.sellIn).To(Equal(1))
			Expect(item.quality).To(Equal(1))

		})
	})

	Describe("given a backstage pass", func() {
		It("increases in quality by 2 if sellIn <=10", func() {
			item := Item{
				name:    "Backstage passes to a TAFKAL80ETC concert",
				sellIn:  10,
				quality: 1,
			}
			err := UpdateQuality([]*Item{&item})
			Expect(err).NotTo(HaveOccurred())
			Expect(item.quality).To(Equal(3))
		})

		It("increases in quality by 3 if sellIn <=5", func() {
			item := Item{
				name:    "Backstage passes to a TAFKAL80ETC concert",
				sellIn:  5,
				quality: 1,
			}
			err := UpdateQuality([]*Item{&item})
			Expect(err).NotTo(HaveOccurred())
			Expect(item.quality).To(Equal(4))
		})

		It("loses all quality if sellIn <=0", func() {
			item := Item{
				name:    "Backstage passes to a TAFKAL80ETC concert",
				sellIn:  0,
				quality: 1,
			}
			err := UpdateQuality([]*Item{&item})
			Expect(err).NotTo(HaveOccurred())
			Expect(item.quality).To(Equal(0))
		})
	})

	Describe("given a conjured item", func() {
		It("degrades in quality by 2", func() {
			item := Item{
				name:    "Conjured item",
				sellIn:  1,
				quality: 2,
			}
			err := UpdateQuality([]*Item{&item})
			Expect(err).NotTo(HaveOccurred())
			Expect(item.quality).To(Equal(0))
		})

		It("degrades to 0 even if odd number", func() {
			item := Item{
				name:    "Conjured item",
				sellIn:  1,
				quality: 1,
			}
			err := UpdateQuality([]*Item{&item})
			Expect(err).NotTo(HaveOccurred())
			Expect(item.quality).To(Equal(0))
		})
	})
})
