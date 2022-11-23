package main

import (
	"github.com/renasami/advents-2022-myjlab/api/controller"
)
func main() {
    router := controller.GetAllRouter()

    router.Run(":8080" )    
}
