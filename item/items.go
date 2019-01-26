package item

import (
	"strings"
)

func UpdateQuality(items []*Item) error {
	for _, item := range items {

		if item.name != "Aged Brie" && item.name != "Backstage passes to a TAFKAL80ETC concert" {
			if item.quality > 0 {
				if strings.HasPrefix(item.name, "Conjured") && item.quality > 1 {
					item.updateQuality(-2)
				} else if item.name != "Sulfuras, Hand of Ragnaros" {
					item.updateQuality(-1)
				}
			}
		} else {
			item.updateQuality(1)
			if item.name == "Backstage passes to a TAFKAL80ETC concert" {
				if item.sellIn < 11 {
					item.updateQuality(1)
				}
				if item.sellIn < 6 {
					item.updateQuality(1)
				}
			}
		}

		if item.name != "Sulfuras, Hand of Ragnaros" {
			item.updateSellIn(-1)
		}

		if item.sellIn < 0 {
			if item.name != "Aged Brie" {
				if item.name != "Backstage passes to a TAFKAL80ETC concert" {
					if item.name != "Sulfuras, Hand of Ragnaros" {
						item.updateQuality(-1)
					}
				} else {
					item.quality = item.quality - item.quality
				}
			} else {
				item.updateQuality(1)
			}
		}
	}
	return nil
}
