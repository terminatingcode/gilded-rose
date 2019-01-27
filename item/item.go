package item

import "fmt"

type Item struct {
	name    string
	sellIn  int
	quality int
}

func New(name string, sellIn, quality int) *Item {
	return &Item{
		name:    name,
		sellIn:  sellIn,
		quality: quality,
	}
}

func (item Item) SellIn() int {
	return item.sellIn
}

func (item Item) Quality() int {
	return item.quality
}

func (item *Item) UpdateSellIn(modifier int) {
	item.sellIn = item.sellIn + modifier
}

func (item *Item) UpdateQuality(modifier int) {
	if item.sellIn <= 0 {
		modifier = modifier * 2
	}
	switch {
	case (item.quality + modifier) > 50:
		item.quality = 50
	case (item.quality + modifier) < 0:
		item.quality = 0
	default:
		item.quality = item.quality + modifier
	}
}

func (item Item) PrettyPrint() {
	fmt.Printf("%-45s %-10d %-10d\n", item.name, item.sellIn, item.quality)
}
