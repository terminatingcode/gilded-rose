package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/terminatingcode/gilded-rose/item"
)

func main() {
	fmt.Println("OMGHAI!")

	var items = []*item.Item{
		item.New("+5 Dexterity Vest", 10, 20),
		item.New("Aged Brie", 2, 0),
		item.New("Elixir of the Mongoose", 5, 7),
		item.New("Sulfuras, Hand of Ragnaros", 0, 80),
		item.New("Sulfuras, Hand of Ragnaros", -1, 80),
		item.New("Backstage passes to a TAFKAL80ETC concert", 15, 20),
		item.New("Backstage passes to a TAFKAL80ETC concert", 10, 49),
		item.New("Backstage passes to a TAFKAL80ETC concert", 5, 49),
		item.New("Conjured Mana Cake", 3, 6), // <-- :O
	}

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

	for day := 0; day < days; day++ {
		fmt.Printf("-------- day %d --------\n", day)
		fmt.Println("name, sellIn, quality")
		for i := 0; i < len(items); i++ {
			fmt.Println(items[i])
		}
		fmt.Println("")
		err := item.UpdateQuality(items)
		if err != nil {
			fmt.Printf("Error on day %d: %s\n", day, err)
		}
	}
}
