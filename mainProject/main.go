package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"test.com/SpellChecker"
	"test.com/YandexSpellChecker"
)

func main() {
	// если проект побольше, необходимо использовать DI
	service := SpellChecker.New(YandexSpellChecker.New("https://speller.yandex.net/services/spellservice.json/checkText?text="))
	router := gin.Default()
	router.GET("/checkText", checkText(service))
	err := router.Run(":8080")
	if err != nil {
		fmt.Println("err: ", err)
	}
}

func checkText(service *SpellChecker.SpellCheckerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := service.Check(c.Query("value"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}
		c.JSON(http.StatusOK, result)
	}
}
