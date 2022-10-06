package main

import (
	"github.com/anfahrul/prb-assistant-api/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/bookcontrol", usecase.InsertBookControl)
	router.GET("/bookcontrol/:medicalRecord", usecase.GetBookControl)
	router.PUT("/bookcontrol/:medicalRecord", usecase.UpdateBook)
	router.GET("/pharmacy", usecase.GetAllPharmacy)
	router.POST("/pharmacy/:medicalRecord", usecase.InsertReceipe)
	router.GET("/pharmacy/recipe", usecase.GetRecipe)
	router.POST("/pharmacy/recipe/:recipeId", usecase.InsertMedicine)
	router.PUT("/pharmacy/:recipeId", usecase.UpdatePharmacy)

	router.Run(":8080")
}
