package product

import "github.com/terminatingcode/gilded-rose/item"

type Variant interface {
	UpdateItem(*item.Item)
}

type DefaultVariant struct{}

func (d DefaultVariant) UpdateItem(item *item.Item) {
	item.UpdateQuality(-1)
	item.UpdateSellIn(-1)
}

type ConjuredVariant struct{}

func (c ConjuredVariant) UpdateItem(item *item.Item) {
	item.UpdateQuality(-2)
	item.UpdateSellIn(-1)
}

type AgedBrieVariant struct{}

func (b AgedBrieVariant) UpdateItem(item *item.Item) {
	item.UpdateQuality(1)
	item.UpdateSellIn(-1)
}

type BackStagePassVariant struct{}

func (bs BackStagePassVariant) UpdateItem(item *item.Item) {
	switch s := item.SellIn(); {
	case s <= 5:
		item.UpdateQuality(3)
	case s <= 10:
		item.UpdateQuality(2)
	default:
		item.UpdateQuality(1)
	}
	item.UpdateSellIn(-1)
}

type SulfurasVariant struct{}

func (s SulfurasVariant) UpdateItem(item *item.Item) {}
