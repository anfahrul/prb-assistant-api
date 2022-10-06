package usecase

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/anfahrul/prb-assistant-api/database"
	"github.com/anfahrul/prb-assistant-api/entity"
	"github.com/anfahrul/prb-assistant-api/repository"
	"github.com/gin-gonic/gin"
)

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

	for i := 1; i <= 3; i++ {
		bookRepository := repository.NewBookRepository(database.GetConnection())
		_, err = bookRepository.InsertBook(ctx, result.MedicalRecord)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
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
	_, err := bookRepository.UpdateBook(ctx, book, int32(medicalRecordInt), int32(bookIdInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"code":    http.StatusCreated,
		"message": "Buku berhasil diupdate",
	})
}
