package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type item struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Quantity int32   `json:"quantity"`
	Price    float64 `json:"price"`
}

var items = []item{
	{ID: "1", Name: "Sugar", Quantity: 10, Price: 50},
	{ID: "2", Name: "Tea", Quantity: 20, Price: 250},
	{ID: "3", Name: "Milk", Quantity: 12, Price: 54},
}

func main() {
	router := gin.Default()
	router.GET("/items", getItems)
	router.POST("/items", postItems)
	router.GET("/item/:id", getItemByID)
	router.Run("localhost:8080")
}

func getItems(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, items)
}

func postItems(c *gin.Context) {
	var newItem item
	if err := c.BindJSON(&newItem); err != nil {
		return
	}
	items = append(items, newItem)
	c.IndentedJSON(http.StatusCreated, newItem)

}

func getItemByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of items, looking for
	// an item whose ID value matches the parameter.
	for _, a := range items {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
}
