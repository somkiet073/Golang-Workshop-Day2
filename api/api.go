package api

import (
	"ginworkshop/model"
	"net/http"

	"log"

	"github.com/gin-gonic/gin"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	// "gopkg.in/mgo.v2/bson"
)

func GetAllPokemon(c *gin.Context) {

	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err.Error())
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	collecttion := session.DB("pokemondb").C("pokemondb")
	pokemons := []model.Pokemon{}
	err = collecttion.Find(nil).All(&pokemons)
	if err != nil {
		log.Fatal(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"result": pokemons})
}

func GetPokemon(c *gin.Context) {
	id := c.Param("id")
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err.Error())
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	collecttion := session.DB("pokemondb").C("pokemondb")
	pokemons := model.Pokemon{}
	// err = collecttion.Find(nil).All(&pokemons)
	err = collecttion.Find(bson.M{"_id": id}).One(&pokemons)
	if err != nil {
		log.Fatal(err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"result": pokemons,
	})
}

func PostPokemon(c *gin.Context) {

	var pokemon model.Pokemon
	c.Bind(&pokemon)
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err.Error())
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	collecttion := session.DB("pokemondb").C("pokemondb")
	// pokemon := model.Pokemon{
	// 	ID:      "0002",
	// 	Name:    "pikaju",
	// 	Element: "222",
	// 	Weight:  "12",
	// }
	err = collecttion.Insert(pokemon)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdatePokemon(c *gin.Context) {
	id := c.Param("id")
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err.Error())
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	collecttion := session.DB("pokemondb").C("pokemondb")
	update := bson.M{
		"$set": bson.M{
			"weight": "10",
		},
	}
	err = collecttion.Update(bson.M{"_id": id}, update)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Success",
	})
}

func DelPokemon(c *gin.Context) {
	id := c.Param("id")
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err.Error())
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	collecttion := session.DB("pokemondb").C("pokemondb")
	err = collecttion.Remove(bson.M{"_id": id})
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Success",
	})
}
