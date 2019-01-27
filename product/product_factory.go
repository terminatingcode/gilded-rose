package product

import (
	"strings"

	"github.com/terminatingcode/gilded-rose/item"
)

func NewProductFactory(name string, sellIn, quality int) Product {
	product := Product{Item: item.New(name, sellIn, quality)}
	switch {
	case strings.Contains(name, "Sulfuras"):
		product.Variant = SulfurasVariant{}
	case strings.Contains(name, "Backstage passes"):
		product.Variant = BackStagePassVariant{}
	case strings.Contains(name, "Aged Brie"):
		product.Variant = AgedBrieVariant{}
	case strings.HasPrefix(name, "Conjured"):
		product.Variant = ConjuredVariant{}
	default:
		product.Variant = DefaultVariant{}
	}
	return product
}
