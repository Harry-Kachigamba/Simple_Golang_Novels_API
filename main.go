// package declaration
package main

//importing http_protocol and gin web framework package
import (
	"errors"
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

// method to get novel by id
func novelById(c *gin.Context) {
	id := c.Param("id")
	novel, err := getNovelById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Novel not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, novel)
}

// method to checkout a book
func checkoutNovel(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
	}

	novel, err := getNovelById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Novel not found."})
		return
	}

	if novel.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Novel not available."})
		return
	}

	novel.Quantity -= 1
	c.IndentedJSON(http.StatusOK, novel)
}

// method to return novel
func returnNovel(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	novel, err := getNovelById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Novel not found."})
		return
	}

	novel.Quantity += 1
	c.IndentedJSON(http.StatusOK, novel)
}

// method to loop through a collection and check
func getNovelById(id string) (*novel, error) {
	for i, n := range novels {
		if n.ID == id {
			return &novels[i], nil
		}
	}

	return nil, errors.New("novel not found")
}

// method to add novel
func createNovel(c *gin.Context) {
	var newNovel novel

	if err := c.BindJSON(&newNovel); err != nil {
		return
	}

	novels = append(novels, newNovel)
	c.IndentedJSON(http.StatusCreated, newNovel)
}

// main method
func main() {
	router := gin.Default()
	router.GET("/novels", getNovels)
	router.POST("/novels", createNovel)
	router.GET("/novels/:id", novelById)
	router.PATCH("/checkout", checkoutNovel)
	router.PATCH("/return", returnNovel)
	router.Run("localhost:8080")
}
