package application

import (
	"testing"

	"github.com/Siriayanur/GoConcurrency/db"
	"github.com/Siriayanur/GoConcurrency/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestApplication(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	var testItems = []models.Item{
		{
			Name:       "Pen",
			Price:      10,
			Quantity:   1,
			Type:       "imported",
			Tax:        1.47,
			FinalPrice: 11.47,
		},
		{
			Name:       "wood",
			Price:      150,
			Quantity:   1,
			Type:       "raw",
			Tax:        18.75,
			FinalPrice: 168.75,
		},
		{
			Name:       "furniture",
			Price:      1000,
			Quantity:   10,
			Type:       "manufactured",
			Tax:        15,
			FinalPrice: 10150,
		},
		{
			Name:       "chocolates",
			Price:      2000,
			Quantity:   5,
			Type:       "imported",
			Tax:        295,
			FinalPrice: 11475,
		},
	}

	mockDB := db.NewMockIDB(mockCtrl)
	mockDB.EXPECT().ReadDataFromDB().Times(1).Return(testItems, nil)
	app := NewApp(mockDB)
	app.RunApp()
	for i := range app.UpdateItems {
		require.Equal(t, app.UpdateItems[i].Name, testItems[i].Name)
		require.Equal(t, app.UpdateItems[i].Price, testItems[i].Price)
		require.Equal(t, app.UpdateItems[i].Type, testItems[i].Type)
		require.Equal(t, app.UpdateItems[i].Quantity, testItems[i].Quantity)
		require.Equal(t, app.UpdateItems[i].Tax, testItems[i].Tax)
		require.Equal(t, app.UpdateItems[i].FinalPrice, testItems[i].FinalPrice)

	}
}
