package usecase

import (
	"context"
	"net/http"
	"strconv"

	"github.com/anfahrul/prb-assistant-api/database"
	"github.com/anfahrul/prb-assistant-api/repository"
	"github.com/gin-gonic/gin"
)

func LoginPatient(c *gin.Context) {
	ctx := context.Background()
	medicalRecord := c.Param("medicalRecord")
	medicalRecordInt, _ := strconv.Atoi(medicalRecord)

	patientRepository := repository.NewPatientRepository(database.GetConnection())
	result, err := patientRepository.FindByMedicalRecordNumber(ctx, int32(medicalRecordInt))
	if err != nil {
		if err.Error() == "medicalRecord 1 Not Found" {
			c.JSON(http.StatusNotFound, err.Error())
			return
		} else {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.JSON(200, result)
}
