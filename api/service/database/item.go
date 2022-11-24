package database

import (
	"database/sql"
	"os"

	"github.com/labstack/gommon/log"
	"github.com/renasami/advents-2022-myjlab/api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func AddItem(item models.Item){

	var dsn = os.Getenv("DATABASE_DSN")
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	row, err := db.Raw("INSERT INTO item values (@id,@value,@author)" ,map[string]interface{}{"id": item.Id.String() , "value": item.Value,"author":item.User}).Rows()
	if err != nil {
		log.Error("Error inserting item")
		return
	}

	row.Close()
}

func GetAllItems(user string) {
	var dsn = os.Getenv("DATABASE_DSN")
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.Raw("SELECT * FROM items WHERE user = @user",sql.Named("user", user))
}