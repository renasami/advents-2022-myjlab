package service

import (
	"os"
	"fmt"
)

func Init () => {
	dsn := os.Getenv("DATABASE_DSN")
	fmt.Println("fa")
	
}