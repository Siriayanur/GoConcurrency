package application

import (
	"testing"

	"github.com/Siriayanur/GoConcurrency/db"
	"github.com/Siriayanur/GoConcurrency/models"
)

func CreateTestData() (IApp, []models.Item) {
	dbi := db.NewDBInstance()
	app := NewApp(dbi)
	items := []models.Item{}
	item1 := models.Item{Name: "a", Price: 100, Type: "raw", Tax: 0.0, FinalPrice: 0.0, Quantity: 2}
	item2 := models.Item{Name: "b", Price: 100, Type: "raw", Tax: 0.0, FinalPrice: 0.0, Quantity: 2}
	item3 := models.Item{Name: "c", Price: 100, Type: "raw", Tax: 0.0, FinalPrice: 0.0, Quantity: 2}
	item4 := models.Item{Name: "d", Price: 100, Type: "raw", Tax: 0.0, FinalPrice: 0.0, Quantity: 2}
	items = append(items, item1)
	items = append(items, item2)
	items = append(items, item3)
	items = append(items, item4)
	return app, items
}
func TestApplication(t *testing.T) {
	app, items := CreateTestData()
	app.AddDataToCollection(items)
}
