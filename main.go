package main

import (
	"github.com/anfahrul/prb-assistant-api/middlewares"
	"github.com/anfahrul/prb-assistant-api/usecase"
	"github.com/gin-gonic/gin"
	// "github.com/rs/cors"
)

func main() {
	router := gin.Default()
	router.Use(middlewares.CORS())

	// all
	public := router.Group("/api")
	public.GET("/pharmacy", usecase.GetAllPharmacy)
	public.GET("/patient/:medicalRecord", usecase.LoginPatient)
	public.GET("/bookcontrol/:medicalRecord", usecase.GetBookControl)
	public.GET("/pharmacy/:recipeId", usecase.GetRecipe)
	public.PUT("/pharmacy/:recipeId", usecase.UpdatePharmacy)
	public.POST("/register", usecase.Register)
	public.POST("/login", usecase.Login)

	// nakes
	nakes := router.Group("/api/nakes")
	// nakes.Use(middlewares.JwtAuthMiddleware())
	nakes.GET("/patients", usecase.GetAllPatient)
	nakes.POST("/logout", usecase.Logout)

	// staf
	staff := router.Group("/api/staff")
	// staff.Use(middlewares.JwtAuthMiddleware())
	// staff.Use(middlewares.StaffMiddleware())
	staff.POST("/bookcontrol", usecase.InsertBookControl)
	staff.POST("/pharmacy/recipe/:recipeId", usecase.InsertMedicine)

	// dokter
	doctor := router.Group("/api/doctor")
	// doctor.Use(middlewares.JwtAuthMiddleware())
	// doctor.Use(middlewares.DoctorMiddleware())
	doctor.PUT("/bookcontrol/:medicalRecord", usecase.UpdateBook)

	router.Run()
}
