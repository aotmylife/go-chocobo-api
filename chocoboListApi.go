package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Character struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type ErrorCase struct {
	Message string `json:"message"`
}

var chocoboDetail = []Character{
	{"chocobo", "images/chocobo.jpg"},
	{"Mog", "images/mog.jpg"},
	{"Golem", "images/golem.jpg"},
	{"Goblin", "images/goblin.jpg"},
	{"Black Mage (Black Magician)", "images/blackmage.jpg"},
	{"White Mage", "images/white_magic.jpg"},
	{"Chubby Chocobo", "images/chubby.jpg"},
	{"Behemoth", "images/behemoth.jpg"},
	{"Bahamut", "images/bahamut.jpg"},
	{"Squall Leonhart", "images/squall_leonhart.jpg"},
}

var errorCase = ErrorCase{
	Message: "message",
}

func main() {
	router := gin.Default()

	router.GET("/chocoboApi", getAllChocoboHandler)
	router.POST("/findChocoboApi", findChocoboHandler)
	router.Static("/chocoboApi/images", "./images")

	router.Run(":8080")
}

func getAllChocoboHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"chocoboDetail": chocoboDetail,
	})
}

func findChocoboHandler(c *gin.Context) {
	var req struct {
		Name string `json:"name"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{errorCase.Message: "Invalid request"})
		return
	}

	if req.Name == "" {
		c.JSON(http.StatusOK, chocoboDetail)
		return
	}

	character, found := findCharacterByName(req.Name)
	if found {
		c.JSON(http.StatusOK, character)
	} else {
		c.JSON(http.StatusNotFound, gin.H{errorCase.Message: "Character not found"})
	}
}

func findCharacterByName(name string) (Character, bool) {
	for _, character := range chocoboDetail {
		if strings.EqualFold(character.Name, name) {
			return character, true
		}
	}
	return Character{}, false
}
