package database

import (
	"errors"
	"fmt"
	"os"

	"github.com/labstack/gommon/log"
	"github.com/renasami/advents-2022-myjlab/api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func AddUser(item models.Auth) (string,error){
	var dsn = os.Getenv("DATABASE_DSN")
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Errorf("gorm ERROR %v",err)
		return "gorm ERROR", err
	}
	result,err := GetUserByEmail(item.Email)
	if err != nil {
		fmt.Printf("gorm ERROR %v",err)
	}
	if result != nil {
		return "this email already exists" , nil
	}
	usr := result
	fmt.Println(usr.Email)
	row, err := db.Raw("INSERT INTO users (name, password, email) values (@name,@password,@email);" ,map[string]interface{}{ "name": item.Username,"password":item.Password,"email":item.Email}).Rows()

	if err != nil {
		log.Error("Error inserting item")
		return "insert error", err
	}
	row.Close()
	return "success", nil
}

func GetUserByEmail(email string) (*models.UserDb, error) {
	var user_result models.UserDb
	var dsn = os.Getenv("DATABASE_DSN")
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	result := db.Raw("SELECT * FROM users WHERE email = ?", email).First(&user_result)
	if errors.Is(result.Error, gorm.ErrRecordNotFound){
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &user_result ,nil

}