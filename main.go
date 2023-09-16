package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// structure of each object of type "car"
type car struct {
	Make  string `json:"make"`
	Model string `json:"model"`
	Color string `json:"color"`
	Power string `json:"power"`
}

// array of data of type "car"
var cars = []car{
	{Make: "BMW", Model: "M3 GTR", Color: "Silver and Blue", Power: "444 BHP"},
	{Make: "Lamborghini", Model: "Gallardo", Color: "Orange", Power: "493 BHP"},
	{Make: "McLaren", Model: "F1", Color: "Blue", Power: "618 BHP"},
	{Make: "Mercedes-Benz", Model: "SLR McLaren", Color: "Black", Power: "617 BHP"},
}

// request handler for GET cars
func getCars(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, cars)
}

// request handler for adding new car to cars array
func addNewCar(c *gin.Context) {
	var newCar car

	if err := c.BindJSON(&newCar); err != nil {
		return
	}

	cars = append(cars, newCar)
	c.IndentedJSON(http.StatusCreated, cars)
}

// request handler for fetching specific car based on url path param
func getCarByMake(c *gin.Context) {
	make := c.Param("make")

	for _, car := range cars {
		if car.Make == make {
			c.IndentedJSON(http.StatusOK, car)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Car not found"})
}

func main() {
	//declare router
	router := gin.Default()

	// link the endpoint url with request handler functions
	router.GET("/cars", getCars)
	router.POST("/car", addNewCar)
	router.GET("/cars/:make", getCarByMake)

	// start the app server
	router.Run("localhost:3000")
}
