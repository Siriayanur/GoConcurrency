package models

import (
	"math"

	"github.com/Siriayanur/GoConcurrency/utils"
)

type Item struct {
	Name       string  `gorm:"default:newItem"`
	Price      float64 `gorm:"default:0"`
	Quantity   int     `gorm:"default:1"`
	Type       string
	Tax        float64 `gorm:"default:0.0"`
	FinalPrice float64 `gorm:"default:0.0"`
}

func NewItem(name string, price float64, quantity int, itemType string) Item {
	item := Item{}
	item.Name = name
	item.Price = price
	item.Quantity = quantity
	item.Type = itemType
	return item
}

func (item *Item) CalculateFinalPrice() {
	mrp := 0.0
	item.Tax = RoundFloat(getTax(item.Type, item.Price))
	mrp = float64(item.Quantity) * (item.Price + item.Tax)
	item.FinalPrice = RoundFloat(mrp)
}

func getTax(itemType string, itemPrice float64) float64 {
	tax := 0.0

	switch itemType {
	case "raw":
		tax = utils.BaseTax * itemPrice
	case "imported":
		tax = utils.BaseTax*itemPrice + (0.02 * (itemPrice + utils.BaseTax*itemPrice))
	case "manufactured":
		tax = utils.ImportDuty * itemPrice
		tax += calculateSurcharge(tax, itemPrice)
	}
	return tax
}

func calculateSurcharge(amount float64, itemPrice float64) float64 {
	switch {
	case amount <= 100:
		return utils.SurchargeLevel1
	case amount <= 200:
		return utils.SurchargeLevel2
	case amount > 200:
		return utils.SurchargeLevel3 * (itemPrice + (itemPrice * utils.ImportDuty))
	default:
		return 0
	}
}
func RoundFloat(val float64) float64 {
	return math.Floor(val*100) / 100
}
