package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/renasami/advents-2022-myjlab/api/controller"
)
func main() {
    router := controller.GetAllRouter()
    // godotenv.Load(config.Path("api", ".env"))
    godotenv.Load()
    log.Print("port",os.Getenv("PORT"))
    log.Println("#env",os.Getenv("DATABASE_DSN"))

    router.Run(":8080" )    
}
