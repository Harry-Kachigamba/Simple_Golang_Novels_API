// package declaration
package main

//importing http_protocol and gin web framework package
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// instance blueprint
type novel struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	Publisher     string `json:"publisher"`
	YearOfPublish string `json:"yearofpublish"`
	Quantity      int    `json:"quantity"`
}

// infromation array
var novels = []novel{
	{ID: "1", Title: "Catch-22", Author: "Joseph Helle", Publisher: "Simon & Schester", YearOfPublish: "1961", Quantity: 20},
	{ID: "2", Title: "The Hitchhiker's Guide to the Galaxy", Author: "Douglas Adams", Publisher: "Pan Books", YearOfPublish: "1979", Quantity: 10},
	{ID: "3", Title: "Good Omens", Author: "Neil Gaiman and Terry Pratchett", Publisher: "Gollancz", YearOfPublish: "1990", Quantity: 30},
	{ID: "4", Title: "Bridget Jone's Diary", Author: "Helen Fielding", Publisher: "Picador", YearOfPublish: "1966", Quantity: 20},
	{ID: "5", Title: "A Confederacy of Dunces", Author: "John Kennedy Toole", Publisher: "Louisiana State University Press", YearOfPublish: "1980", Quantity: 10},
}

// method: responds to incoming http requests
func getNovels(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, novels)
}

// main method
func main() {
	router := gin.Default()
	router.GET("/novels", getNovels)
	router.Run("localhost:8080")
}
