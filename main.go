package main

import (
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}


// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
    var newAlbum album

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    // Add the new album to the slice.
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}


func main() {

	fmt.Println(albums)
	fmt.Println(albums[0])
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id",getParticularAlbum)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8081")

}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}


func getParticularAlbum(c *gin.Context) {
	id := c.Param("id")
	//fmt.Println(id)
	for _,a := range albums{
		if a.ID == id{
			i, err := strconv.Atoi(a.ID)
    		if err != nil {
        			panic(err)
   			 }
			c.IndentedJSON(http.StatusOK, albums[i])
            return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})

}