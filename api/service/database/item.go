package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/labstack/gommon/log"
	"github.com/renasami/advents-2022-myjlab/api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func AddItem(item models.Item, author string) bool {

	var dsn = os.Getenv("DATABASE_DSN")
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	fmt.Println(dsn)
	row, err := db.Raw("INSERT INTO items values (@id,@value,@author,@done)", map[string]interface{}{"id": item.Id.String(), "value": item.Value, "author": author, "done": item.Done}).Rows()
	if err != nil {
		log.Error("Error inserting item")
		return false
	}
	return true
	row.Close()
	return false
}

func GetAllItems(user string) *[]models.ItemFromDB {
	var items []models.ItemFromDB
	var dsn = os.Getenv("DATABASE_DSN")
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.Raw("SELECT * FROM items WHERE author = @user", sql.Named("user", user)).Scan(&items)
	fmt.Println(items)
	return &items
}

func DeleteItemWithId(id string) bool {
	var dsn = os.Getenv("DATABASE_DSN")
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	row, err := db.Raw("DELETE FROM items WHERE id = @id", sql.Named("id", id)).Rows()
	if err != nil {
		log.Error("Error inserting item")
		return false
	}
	row.Close()
	return true
}

func EditItem(item models.Item) bool {
	var dsn = os.Getenv("DATABASE_DSN")
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	fmt.Println(dsn)
	row, err := db.Raw("UPDATE items SET value = @value, done = @done where id = @id", map[string]interface{}{"id": item.Id.String(), "value": item.Value, "done": item.Done}).Rows()
	if err != nil {
		log.Error("Error inserting item")
		return false
	}
	return true
	row.Close()
	return false
}
