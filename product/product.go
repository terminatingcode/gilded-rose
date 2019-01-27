package product

import (
	"fmt"

	"github.com/terminatingcode/gilded-rose/item"
)

type Product struct {
	Item    *item.Item
	Variant Variant
}

func (p Product) Update() error {
	if p.Item == nil {
		return fmt.Errorf("Product %v was not initialised with an item", p.Variant)
	}

	p.Variant.UpdateItem(p.Item)
	return nil
}
