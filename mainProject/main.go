package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"test.com/SpellChecker"
	"test.com/Viper"
	"test.com/YandexSpellChecker"
)

func main() {
	config, err := Viper.ConfigViper.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	service := SpellChecker.New(YandexSpellChecker.New(config.YANDEXURL))
	router := gin.Default()
	router.GET("/checkText", checkText(service))
	err = router.Run(":8080")
	if err != nil {
		log.Fatal("err: ", err)
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
