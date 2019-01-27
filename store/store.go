package store

import (
	"fmt"

	"github.com/terminatingcode/gilded-rose/product"
)

type Store struct {
	Products []product.Product
}

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

func (s Store) PrintProducts(day int) {
	fmt.Printf("-------- day %d --------\n", day)
	fmt.Printf("%-45s %-10s %-10s\n", "name", "sellIn", "quality")
	for _, product := range s.Products {
		product.Item.PrettyPrint()
	}
}
