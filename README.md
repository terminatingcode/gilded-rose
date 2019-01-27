# Introduction
Master branch is the translation of [gilded-rose](https://github.com/NotMyself/GildedRose)
by [Emily Bache](https://github.com/emilybache/GildedRose-Refactoring-Kata) before refactoring.

Gilded rose is a practice in refactoring without breaking the existing contract.
Navigate to other branches to see implementation of difference versions.

## Minimum Implementation

`git checkout minimum-implementation`

This branch contacts the minimal amount of code needed to implement the new item required by the system.
This has the lowest cost of time in the short term.

## Factory Design

`git checkout product-factory`

This design follows a [factory design pattern](https://www.sohamkamani.com/blog/golang/2018-06-20-golang-factory-patterns/)
to create a new struct `Product`:
```
type Product struct {
	Item    *item.Item
	Variant Variant
}
```

Where each Variant can contain it's own logic on how to update an item:
```
type Variant interface {
	UpdateItem(*item.Item)
}
```

Each product can be created using a factory method which switches on the products name:

```
NewProductFactory(name string, sellIn, quality int) Product {
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
```

Finally, a `Store` can keep a slice of Products and call their update methods for as many days as needed to update the inventory:
```
func (s Store) UpdateInventory(days int) error {
	for i := 0; i < days; i++ {
		s.PrintProducts(i)
		for _, product := range s.Products {
			err := product.Update()
			if err != nil {
				return err
			}
		}
	}
	return nil
}
```

