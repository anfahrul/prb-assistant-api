package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/anfahrul/prb-assistant-api/database"
	"github.com/anfahrul/prb-assistant-api/entity"
	"github.com/anfahrul/prb-assistant-api/repository"
	"github.com/gin-gonic/gin"
)

func GetAllPatient(c *gin.Context) {
	ctx := context.Background()

	patientRepository := repository.NewPatientRepository(database.GetConnection())
	result, err := patientRepository.GetAllPatient(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(200, result)
}

func InsertBookControl(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	ctx := context.Background()
	bodyReq, _ := ioutil.ReadAll(c.Request.Body)
	var patient entity.Patient
	json.Unmarshal(bodyReq, &patient)

	patientRepository := repository.NewPatientRepository(database.GetConnection())
	result, err := patientRepository.Insert(ctx, patient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	currentTime := time.Now().UnixMilli()
	milliOfMonth := 2629800000
	for i := 1; i <= 3; i++ {
		bookRepository := repository.NewBookRepository(database.GetConnection())
		time := (milliOfMonth * i) + int(currentTime)
		_, err = bookRepository.InsertBook(ctx, result.MedicalRecord, time)
		if err != nil {
			fmt.Println("disini")
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
	}

	var recipe entity.Recipe
	recipeRepository := repository.NewRecipeRepository(database.GetConnection())
	_, err = recipeRepository.InsertRecipe(ctx, int64(result.MedicalRecord), recipe)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"code":    http.StatusCreated,
		"message": "create succesfull",
	})
}

func GetBookControl(c *gin.Context) {
	ctx := context.Background()
	medicalRecord := c.Param("medicalRecord")
	medicalRecordInt, _ := strconv.Atoi(medicalRecord)

	patientRepository := repository.NewPatientRepository(database.GetConnection())
	result, err := patientRepository.FindByMedicalRecordNumber(ctx, int32(medicalRecordInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	bookRepository := repository.NewBookRepository(database.GetConnection())
	bookResult, err := bookRepository.FindBookByMedicalRecordNumber(ctx, result.MedicalRecord)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"patientData": result,
		"books":       bookResult,
	})
}

func UpdateBook(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	ctx := context.Background()
	bodyReq, _ := ioutil.ReadAll(c.Request.Body)
	var book entity.Book
	json.Unmarshal(bodyReq, &book)

	medicalRecord := c.Param("medicalRecord")
	medicalRecordInt, _ := strconv.Atoi(medicalRecord)
	bookId := c.Query("bookId")
	bookIdInt, _ := strconv.Atoi(bookId)

	bookRepository := repository.NewBookRepository(database.GetConnection())
	_, err := bookRepository.FindBookByMedicalRecordNumber(ctx, int64(medicalRecordInt))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": err.Error(),
		})
		return
	}

	err = bookRepository.FindBookById(ctx, int32(bookIdInt), int64(medicalRecordInt))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": err.Error(),
		})
		return
	}

	_, err = bookRepository.UpdateBook(ctx, book, int64(medicalRecordInt), int32(bookIdInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code":    http.StatusCreated,
		"message": "Buku berhasil diupdate",
	})
}
