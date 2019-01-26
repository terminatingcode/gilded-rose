package item

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

func (item *Item) updateSellIn(modifier int) {
	item.sellIn = item.sellIn + modifier
}

func (item *Item) updateQuality(modifier int) {
	if item.quality < 50 {
		item.quality = item.quality + modifier
	}
}
