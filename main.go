package main

import (
	"ginworkshop/api"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/GetAllPokemon", api.GetAllPokemon)
	r.GET("/GetPokemon/:id", api.GetPokemon)
	r.POST("/", api.PostPokemon)
	r.PUT("/:id", api.UpdatePokemon)
	r.DELETE("/:id", api.DelPokemon)

	r.Run(":8111")
}
