package main

import (
	"os"

	"github.com/renasami/advents-2022-myjlab/api/controller"
)
func main() {
    router := controller.GetAllRouter()
    port := os.Getenv("PORT")
    router.Run(":" + port)    
}
