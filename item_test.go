package main

import (
	"testing"
)

var items = []struct {
	name            string
	description     string
	days            int
	quality         int
	expectedDays    int
	expectedQuality int
}{
	{"normal", "before sell date", 5, 10, 4, 9},
	{"normal", "on sell date", 0, 10, -1, 8},
	{"normal", "after sell date", -10, 10, -11, 8},
	{"normal", "of zero quality", 5, 0, 4, 0},
}

var bries = []struct {
	name            string
	description     string
	days            int
	quality         int
	expectedDays    int
	expectedQuality int
}{
	{"Aged Brie", "before sell date with max quality", 5, 50, 4, 50},
	{"Aged Brie", "on sell date", 0, 10, -1, 12},
	{"Aged Brie", "on sell date near max quality", 0, 49, -1, 50},
	{"Aged Brie", "on sell date with max quality", 0, 50, -1, 50},
	{"Aged Brie", "after sell date", -10, 10, -11, 12},
	{"Aged Brie", "after sell date with max quality", -10, 50, -11, 50},
}

var legendaries = []struct {
	name            string
	description     string
	quality         int
	expectedQuality int
}{
	{"Sulfuras, Hand of Ragnaros", "before sell date", 80, 80},
}

var backstagePasses = []struct {
	name            string
	description     string
	days            int
	quality         int
	expectedDays    int
	expectedQuality int
}{
	{"Backstage passes to a TAFKAL80ETC concert", "long before sell date", 11, 10, 10, 11},
	{"Backstage passes to a TAFKAL80ETC concert", "long before sell date at max quality", 11, 50, 10, 50},
	{"Backstage passes to a TAFKAL80ETC concert", "medium close to sell date (upper bound)", 10, 10, 9, 12},
	{"Backstage passes to a TAFKAL80ETC concert", "medium close to sell date (upper bound) at max quality", 10, 50, 9, 50},
	{"Backstage passes to a TAFKAL80ETC concert", "medium close to sell date (lower bound)", 6, 10, 5, 12},
	{"Backstage passes to a TAFKAL80ETC concert", "medium close to sell date (lower bound) at max quality", 6, 50, 5, 50},
	{"Backstage passes to a TAFKAL80ETC concert", "very close to sell date (upper bound)", 5, 10, 4, 13},
	{"Backstage passes to a TAFKAL80ETC concert", "very close to sell date (upper bound) at max quality", 5, 50, 4, 50},
	{"Backstage passes to a TAFKAL80ETC concert", "very close to sell date (lower bound)", 1, 10, 0, 13},
	{"Backstage passes to a TAFKAL80ETC concert", "very close to sell date (lower bound) at max quality", 1, 50, 0, 50},
	{"Backstage passes to a TAFKAL80ETC concert", "on sell date", 0, 10, -1, 0},
	{"Backstage passes to a TAFKAL80ETC concert", "after sell date", -10, 10, -11, 0},
}

var conjuredItems = []struct {
	name            string
	description     string
	days            int
	quality         int
	expectedDays    int
	expectedQuality int
}{

	{"Conjured Mana Cake", "before sell date", 5, 10, 4, 8},
	{"Conjured Mana Cake", "before sell date at zero quality", 5, 0, 4, 0},
	{"Conjured Mana Cake", "on sell date", 0, 10, -1, 6},
	{"Conjured Mana Cake", "on sell date at zero quality", 0, 0, -1, 0},
	{"Conjured Mana Cake", "after sell date", -10, 10, -11, 6},
	{"Conjured Mana Cake", "after sell date at zero quality", -10, 0, -11, 0},
}

func TestGildedRoseItems(t *testing.T) {
	for _, tt := range items {
		item := Item{name: tt.name, days: tt.days, quality: tt.quality}
		item.updateQuality()
		item.age()
		if item.quality != tt.expectedQuality {
			t.Errorf("\n%s %s\nexpected quality: %d\nactual quality: %d", tt.name, tt.description, tt.expectedQuality, item.quality)
		}
		if item.days != tt.expectedDays {
			t.Errorf("\n%s %s\nexpected days left: %d\nactual days left: %d", tt.name, tt.description, tt.expectedDays, item.days)
		}
	}
}

func TestGildedRoseBries(t *testing.T) {
	for _, tt := range bries {
		item := CheeseItem{name: tt.name, days: tt.days, quality: tt.quality}
		item.updateQuality()
		item.age()
		if item.quality != tt.expectedQuality {
			t.Errorf("\n%s %s\nexpected quality: %d\nactual quality: %d", tt.name, tt.description, tt.expectedQuality, item.quality)
		}
		if item.days != tt.expectedDays {
			t.Errorf("\n%s %s\nexpected days left: %d\nactual days left: %d", tt.name, tt.description, tt.expectedDays, item.days)
		}
	}
}

func TestGildedRoseLegendaryItems(t *testing.T) {
	for _, tt := range legendaries {
		item := LegendaryItem{name: tt.name, quality: tt.quality}
		if item.quality != tt.expectedQuality {
			t.Errorf("\n%s %s\nexpected quality: %d\nactual quality: %d", tt.name, tt.description, tt.expectedQuality, item.quality)
		}
	}
}

func TestGildedRoseBackstagePasses(t *testing.T) {
	for _, tt := range backstagePasses {
		item := BackstagePassItem{name: tt.name, days: tt.days, quality: tt.quality}
		item.updateQuality()
		item.age()
		if item.quality != tt.expectedQuality {
			t.Errorf("\n%s %s\nexpected quality: %d\nactual quality: %d", tt.name, tt.description, tt.expectedQuality, item.quality)
		}
		if item.days != tt.expectedDays {
			t.Errorf("\n%s %s\nexpected days left: %d\nactual days left: %d", tt.name, tt.description, tt.expectedDays, item.days)
		}
	}
}

func TestGildedRoseConjuredItems(t *testing.T) {
	for _, tt := range conjuredItems {
		item := ConjuredItem{name: tt.name, days: tt.days, quality: tt.quality}
		item.updateQuality()
		item.age()
		if item.quality != tt.expectedQuality {
			t.Errorf("\n%s %s\nexpected quality: %d\nactual quality: %d", tt.name, tt.description, tt.expectedQuality, item.quality)
		}
		if item.days != tt.expectedDays {
			t.Errorf("\n%s %s\nexpected days left: %d\nactual days left: %d", tt.name, tt.description, tt.expectedDays, item.days)
		}
	}
}
