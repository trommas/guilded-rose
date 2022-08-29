package main

import "fmt"

type Inventory struct {
	items              []Item
	legendaryItems     []LegendaryItem
	cheesesItems       []CheeseItem
	backstagePassItems []BackstagePassItem
	conjuredItems      []ConjuredItem
}

func (i *Inventory) addItem(item Item) *Inventory {
	i.items = append(i.items, item)
	return i
}

func (i *Inventory) addLegendaryItem(item LegendaryItem) *Inventory {
	i.legendaryItems = append(i.legendaryItems, item)
	return i
}

func (i *Inventory) addCheeseItem(item CheeseItem) *Inventory {
	i.cheesesItems = append(i.cheesesItems, item)
	return i
}

func (i *Inventory) addBackstagePassItem(item BackstagePassItem) *Inventory {
	i.backstagePassItems = append(i.backstagePassItems, item)
	return i
}

func (i *Inventory) addConjuredItem(item ConjuredItem) *Inventory {
	i.conjuredItems = append(i.conjuredItems, item)
	return i
}

func (i Inventory) print() {
	fmt.Println("Inventory:")
	for _, item := range i.items {
		fmt.Printf("Item: %s \n", item.name)
		fmt.Printf("\t Days left: %d \t Quality: %d \n", item.days, item.quality)
	}
	for _, item := range i.legendaryItems {
		fmt.Printf("Legendary Item: %s \n", item.name)
		fmt.Printf("\t Quality: %d \n", item.quality)
	}
	for _, item := range i.cheesesItems {
		fmt.Printf("Cheese: %s \n", item.name)
		fmt.Printf("\t Days left: %d \t Quality: %d \n", item.days, item.quality)
	}
	for _, item := range i.backstagePassItems {
		fmt.Printf("Backstage Pass: %s \n", item.name)
		fmt.Printf("\t Days left: %d \t Quality: %d \n", item.days, item.quality)
	}
	for _, item := range i.conjuredItems {
		fmt.Printf("Conjured Item: %s \n", item.name)
		fmt.Printf("\t Days left: %d \t Quality: %d \n", item.days, item.quality)
	}
}

// AgeItems ages the items by a day
func (inv Inventory) ageItems() {
	for i := range inv.items {
		inv.items[i].age()
	}
	for i := range inv.cheesesItems {
		inv.cheesesItems[i].age()
	}
	for i := range inv.backstagePassItems {
		inv.backstagePassItems[i].age()
	}
	for i := range inv.backstagePassItems {
		inv.backstagePassItems[i].age()
	}
	for i := range inv.conjuredItems {
		inv.conjuredItems[i].age()
	}
}

// UpdateQuality updates the quality of the items
func (inv Inventory) updateItemsQuality() {
	for i := range inv.items {
		inv.items[i].updateQuality()
	}
	for i := range inv.cheesesItems {
		inv.cheesesItems[i].updateQuality()
	}
	for i := range inv.backstagePassItems {
		inv.backstagePassItems[i].updateQuality()
	}
	for i := range inv.conjuredItems {
		inv.conjuredItems[i].updateQuality()
	}
}
