package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Doc for Go struct tags (or struct field tags): https://go.dev/ref/spec#Struct_types
// Doc for JSON specific can be found in the book: "The Go Programming Language" in the sub-section: 4.5 JSON, on page 109,
//
//	Which tells us that the json key controls the behaviour of the encoding/json package, and other encoding packages follow this convention.
//	So if we look at the official Go doc for the encoding/json package (https://pkg.go.dev/encoding/json) we get a little more information.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "cats", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) { // Try with context as being a copy rather than a pointer to the context
	// Doc: https://pkg.go.dev/github.com/gin-gonic/gin#Context.IndentedJSON
	// Amongst other things, IndentedJSON serializes the given struct (or in our case a slice of album struct instances) as pretty JSON (indented + endlines) into the response body.
	if gin.Mode() == gin.DebugMode {
		c.IndentedJSON(http.StatusOK, albums)
	} else {
		c.JSON(http.StatusOK, albums)
	}
}

// The user should not be adding the id, as they could add duplicate values, this is now handled programmatically.
// The payload should be validated before processing for at least two reasons.
//
//	First security
//	Second : Example: If an object id is not a string, the current code can't handle that (returns a 400)
func postAlbums(c *gin.Context) {
	var newAlbum album
	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	// Find out what the highest ID is
	var highestID int
	for _, album := range albums {
		id, err := strconv.Atoi(album.ID)
		if err == nil && id > highestID {
			highestID = id
		}
	}
	// Assign the highest number + 1 to album
	newAlbum.ID = strconv.Itoa(highestID + 1)
	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	if gin.Mode() == gin.DebugMode {
		c.IndentedJSON(http.StatusCreated, newAlbum)
	}
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	gin.SetMode(gin.DebugMode)
	//gin.SetMode(gin.ReleaseMode)

	// If you use gin.New() instead of gin.Default(), and you want to add logger and recovery middleware, you need to do it manually.
	//   gin.New() obviously also means the logged warnings (that aren't actually warnings, disappear).
	router := gin.Default()
	//router := gin.New()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums", postAlbums)

	// Doc: https://pkg.go.dev/github.com/gin-gonic/gin#Engine.Run
	//   Blocks the calling routine indefinitely unless an error happens.
	router.Run("localhost:8080")
}
