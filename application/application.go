package application

import (
	"fmt"
	"os"
	"sync"

	"github.com/Siriayanur/GoConcurrency/db"
	"github.com/Siriayanur/GoConcurrency/models"
)

type App struct {
	wg            sync.WaitGroup
	mutex         sync.Mutex
	dbInstance    *db.DB
	retrieveItems []models.Item
	updateItems   []models.Item
}

func NewApp(di *db.DB) *App {
	app := App{dbInstance: di}
	return &app
}
func (app *App) RunApp() {
	// wait till all go routines are executed
	app.wg.Add(1)

	data, err := app.dbInstance.ReadDataFromDB()
	if err != nil {
		fmt.Println("DB Error :: ", err)
		os.Exit(1)
	}
	addItemsToCollectionChannel := app.AddDataToCollection(data)
	useItemsToCalcTaxChannel := app.Calculate(addItemsToCollectionChannel)
	updateItemsToCollectionChannel := app.UpdateItemToCollection(useItemsToCalcTaxChannel)
	app.DisplayItems(updateItemsToCollectionChannel)
	app.wg.Wait()
}

// Add items to collection(retrieveItems) and send to channel store1.
func (app *App) AddDataToCollection(items []models.Item) chan models.Item {
	store := make(chan models.Item)
	go func() {
		defer close(store)
		for _, val := range items {
			store <- val
			app.mutex.Lock()
			app.retrieveItems = append(app.retrieveItems, val)
			app.mutex.Unlock()
		}
	}()
	return store
}

// Retrieve the items from store1, calculate Tax for each item, send to channel store2.
func (app *App) Calculate(retrieve chan models.Item) chan models.Item {
	store := make(chan models.Item)
	go func() {
		defer close(store)
		for val := range retrieve {
			val.CalculateFinalPrice()
			store <- val
		}
	}()
	return store
}

// Retrieve items from store2, populate updateItems collection, send to channel store3.
func (app *App) UpdateItemToCollection(retrieve chan models.Item) chan models.Item {
	store := make(chan models.Item)
	go func() {
		defer close(store)
		for val := range retrieve {
			// ensure updateItems is not used by any other routine
			app.mutex.Lock()
			app.updateItems = append(app.updateItems, val)
			app.mutex.Unlock()
			store <- val
		}

	}()
	return store
}

// Retrieve items from channel store3 to console to user.
func (app *App) DisplayItems(retrieve chan models.Item) {
	go func() {
		for val := range retrieve {
			fmt.Printf("NAME :: %s | QUANTITY :: %d | TAX :: %v | TOTAL :: %v\n", val.Name, val.Quantity, val.Tax, val.FinalPrice)
		}
		app.wg.Done()
	}()
	// app.wg.Done()
}
