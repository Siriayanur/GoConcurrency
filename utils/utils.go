package utils

import "math"

const (
	DefaultItemName     string  = "Item"
	DefaultItemQuantity int     = 1
	DefaultPrice        float64 = 0.0
	BaseTax                     = 0.125
	ImportDuty                  = 0.01
	SurchargeLevel1             = 5.0
	SurchargeLevel2             = 10.0
	SurchargeLevel3             = 0.05
)

func RoundFloat(val float64) float64 {
	return math.Floor(val*100) / 100
}
