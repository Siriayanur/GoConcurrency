package db

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/Siriayanur/GoConcurrency/config"
	"github.com/Siriayanur/GoConcurrency/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

type IDB interface {
	InitDB() error
	AddDataToDB() error
	ReadDataFromDB() ([]models.Item, error)
}

func (d *DB) InitDB() error {
	conf := config.LoadConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.Database.User, conf.Database.Password, conf.Database.Host, conf.Database.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}
	d.db = db
	if !db.Migrator().HasTable(&models.Item{}) {
		err := d.db.AutoMigrate(&models.Item{})
		if err != nil {
			return err
		}
		err = d.AddDataToDB()
		if err != nil {
			return err
		}
	}
	return nil
}

// populate item data.
func (d *DB) AddDataToDB() error {
	data, err := ReadFileData()
	if err != nil {
		return err
	}
	for _, val := range data {
		d.db.Create(&val)
	}
	return nil
}
func (d *DB) ReadDataFromDB() ([]models.Item, error) {
	var items []models.Item
	data := d.db.Find(&items)

	if data.Error != nil {
		return nil, data.Error
	}
	return items, nil
}
func ReadFileData() ([]models.Item, error) {
	filePointer, err := os.Open("itemData.json")
	if err != nil {
		return nil, err
	}
	defer filePointer.Close()
	marshalData, err := io.ReadAll(filePointer)
	if err != nil {
		return nil, err
	}

	var unmarshalData []models.Item
	if len(marshalData) == 0 {
		return unmarshalData, nil
	}
	err = json.Unmarshal(marshalData, &unmarshalData)
	if err != nil {
		return nil, err
	}
	return unmarshalData, nil
}
func NewDBInstance() IDB {
	db := DB{}
	err := db.InitDB()

	if err != nil {
		fmt.Println("DB Error :: Couldn't create")
		os.Exit(1)
	}
	return &db
}
