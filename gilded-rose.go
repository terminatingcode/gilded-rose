package main

import (
	"strings"
)

// if sellIn 0, quality -x2
// quality never <0
// aged brie increases in quality as time passes
// max quality 50
// sulfuras no sellIn or degradein quality
// backstage passes: if sellIn <=10 +2, if sellIn <=5 +3, if sellIn==0 0
type Item struct {
	name    string
	sellIn  int
	quality int
}

func UpdateQuality(items []*Item) error {
	for i := 0; i < len(items); i++ {

		if items[i].name != "Aged Brie" && items[i].name != "Backstage passes to a TAFKAL80ETC concert" {
			if items[i].quality > 0 {
				if strings.HasPrefix(items[i].name, "Conjured") && items[i].quality > 1 {
					items[i].quality = items[i].quality - 2
				} else if items[i].name != "Sulfuras, Hand of Ragnaros" {
					items[i].quality = items[i].quality - 1
				}
			}
		} else {
			if items[i].quality < 50 {
				items[i].quality = items[i].quality + 1
				if items[i].name == "Backstage passes to a TAFKAL80ETC concert" {
					if items[i].sellIn < 11 {
						if items[i].quality < 50 {
							items[i].quality = items[i].quality + 1
						}
					}
					if items[i].sellIn < 6 {
						if items[i].quality < 50 {
							items[i].quality = items[i].quality + 1
						}
					}
				}
			}
		}

		if items[i].name != "Sulfuras, Hand of Ragnaros" {
			items[i].sellIn = items[i].sellIn - 1
		}

		if items[i].sellIn < 0 {
			if items[i].name != "Aged Brie" {
				if items[i].name != "Backstage passes to a TAFKAL80ETC concert" {
					if items[i].quality > 0 {
						if items[i].name != "Sulfuras, Hand of Ragnaros" {
							items[i].quality = items[i].quality - 1
						}
					}
				} else {
					items[i].quality = items[i].quality - items[i].quality
				}
			} else {
				if items[i].quality < 50 {
					items[i].quality = items[i].quality + 1
				}
			}
		}
	}
	return nil
}
