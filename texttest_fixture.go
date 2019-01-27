package main

import (
	"fmt"
	"os"
	"strconv"

	p "github.com/terminatingcode/gilded-rose/product"
	s "github.com/terminatingcode/gilded-rose/store"
)

func main() {
	fmt.Println("OMGHAI!")

	var products = []p.Product{
		p.NewProductFactory("+5 Dexterity Vest", 10, 20),
		p.NewProductFactory("Aged Brie", 2, 0),
		p.NewProductFactory("Elixir of the Mongoose", 5, 7),
		p.NewProductFactory("Sulfuras, Hand of Ragnaros", 0, 80),
		p.NewProductFactory("Sulfuras, Hand of Ragnaros", -1, 80),
		p.NewProductFactory("Backstage passes to a TAFKAL80ETC concert", 15, 20),
		p.NewProductFactory("Backstage passes to a TAFKAL80ETC concert", 10, 49),
		p.NewProductFactory("Backstage passes to a TAFKAL80ETC concert", 5, 49),
		p.NewProductFactory("Conjured Mana Cake", 3, 6), // <-- :O
	}

	store := s.Store{products}

	days := 2
	var err error
	if len(os.Args) > 1 {
		days, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		days++
	}

	err = store.UpdateInventory(days)
	if err != nil {
		fmt.Printf("Error updating inventory: %s\n", err)
	}
}
