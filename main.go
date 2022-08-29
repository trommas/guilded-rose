package main

func main() {
	i := &Inventory{}

	i.addItem(Item{name: "+5 Dexterity Vest", days: 10, quality: 20})
	i.addCheeseItem(CheeseItem{name: "Aged Brie", days: 2, quality: 0})
	i.addItem(Item{name: "Elixir of the Mongoose", days: 5, quality: 7})
	i.addLegendaryItem(LegendaryItem{name: "Sulfuras, Hand of Ragnaros", quality: 80})
	i.addBackstagePassItem(BackstagePassItem{name: "Backstage passes to a TAFKAL80ETC concert", days: 2, quality: 50})
	i.addConjuredItem(ConjuredItem{name: "Love potion", days: 5, quality: 12})

	i.print()

	i.ageItems()
	i.updateItemsQuality()

	i.print()
}
