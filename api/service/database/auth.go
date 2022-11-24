package database

import (
	"os"

	"github.com/labstack/gommon/log"
	"github.com/renasami/advents-2022-myjlab/api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func AddUser(item models.Auth){
	var dsn = os.Getenv("DATABASE_DSN")
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Errorf("gorm ERROR %v",err)
		return
	}
	result,err := GetUserByEmail(item.Email)
	log.Print("result",result)
	log.Printf("err",err)
	return
	row, err := db.Raw("INSERT INTO users (name, password, email) values (@name,@password,@email);" ,map[string]interface{}{ "name": item.Username,"password":item.Password,"email":item.Email}).Rows()
	log.Print("row",row)
	if err != nil {
		log.Error("Error inserting item")
		return
	}

	row.Close()
}

func GetUserByEmail(email string) (*models.UserDb, error) {
	var user_result models.UserDb
	var dsn = os.Getenv("DATABASE_DSN")
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	result := db.Raw("SELECT * FROM users WHERE email = ?", email).First(&user_result)
	log.Print("row",result)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user_result ,nil

}