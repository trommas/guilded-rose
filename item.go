package main

type Artifact interface {
	age()
	updateQuality()
}

// Item describes an item sold by the Gilded Rose Inn. NOT ALLOWED TO MODIFY
type Item struct {
	name    string
	days    int
	quality int
}

type LegendaryItem struct {
	name    string
	quality int
}

type CheeseItem struct {
	name    string
	days    int
	quality int
}

type BackstagePassItem struct {
	name    string
	days    int
	quality int
}

type ConjuredItem struct {
	name    string
	days    int
	quality int
}

func (item *CheeseItem) updateQuality() {
	if item.days > 0 {
		item.quality++
	} else {
		item.quality = item.quality + 2
	}
	item.quality = sanitizeQuality(item.quality)
}

func (item *BackstagePassItem) updateQuality() {
	if item.days > 10 {
		item.quality++
	} else if item.days > 5 {
		item.quality = item.quality + 2
	} else if item.days > 0 {
		item.quality = item.quality + 3
	} else {
		item.quality = 0
	}
	item.quality = sanitizeQuality(item.quality)
}

func (item *Item) updateQuality() {
	if item.days > 0 {
		item.quality--
	} else {
		item.quality = item.quality - 2
	}
	item.quality = sanitizeQuality(item.quality)
}

func (item *ConjuredItem) updateQuality() {
	if item.days > 0 {
		item.quality = item.quality - 2
	} else {
		item.quality = item.quality - 4
	}
	item.quality = sanitizeQuality(item.quality)
}

func sanitizeQuality(q int) int {
	if q > 50 {
		return 50
	}
	if q < 0 {
		return 0
	}
	return q
}

func (i *Item) age() {
	i.days--
}

func (i *CheeseItem) age() {
	i.days--
}

func (i *BackstagePassItem) age() {
	i.days--
}

func (i *ConjuredItem) age() {
	i.days--
}
