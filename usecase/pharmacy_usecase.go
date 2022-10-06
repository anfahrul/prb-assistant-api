package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/anfahrul/prb-assistant-api/database"
	"github.com/anfahrul/prb-assistant-api/entity"
	"github.com/anfahrul/prb-assistant-api/repository"
	"github.com/gin-gonic/gin"
)

func InsertReceipe(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	ctx := context.Background()
	medicalRecord := c.Param("medicalRecord")
	medicalRecordInt, _ := strconv.Atoi(medicalRecord)
	var recipe entity.Recipe

	recipeRepository := repository.NewRecipeRepository(database.GetConnection())
	_, err := recipeRepository.InsertRecipe(ctx, int64(medicalRecordInt), recipe)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"code":    http.StatusCreated,
		"message": "create succesfull",
	})
}

func GetRecipe(c *gin.Context) {
	ctx := context.Background()
	recipeId := c.Query("recipeId")
	recipeIdInt, _ := strconv.Atoi(recipeId)
	fmt.Println("RESPID", recipeIdInt)

	recipeRepository := repository.NewRecipeRepository(database.GetConnection())
	recipe, err := recipeRepository.FindByRecipeId(ctx, int64(recipeIdInt))
	if err != nil {
		fmt.Println("RECEIPE", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	patientRepository := repository.NewPatientRepository(database.GetConnection())
	patient, err := patientRepository.FindByMedicalRecordNumber(ctx, recipe.MedicalRecord)
	if err != nil {
		fmt.Println("PATIENT", err.Error())
		fmt.Println("PATIENT", recipe)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	pharmacyRepository := repository.NewPharmacyRepository(database.GetConnection())
	pharmacy, err := pharmacyRepository.FindByPharmacyId(ctx, int64(recipe.PharmacyId))
	if err != nil {
		if pharmacy.PharmacyId == 0 {
		} else {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
	}

	medicineRepository := repository.NewMedicineRepository(database.GetConnection())
	medicine, err := medicineRepository.FindByRecipeId(ctx, int64(recipeIdInt))
	if err != nil {
		fmt.Println("MEDICINE", err.Error())
		if medicine[0].RecipeId == 0 {
		} else {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.JSON(200, gin.H{
		"recipeId":    recipeId,
		"patient":     patient,
		"claimStatus": recipe.ClaimStatus,
		"pharmacy":    pharmacy,
		"medicine":    medicine,
	})
}

func InsertMedicine(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	ctx := context.Background()
	recipeId := c.Param("recipeId")
	recipeIdInt, _ := strconv.Atoi(recipeId)
	bodyReq, _ := ioutil.ReadAll(c.Request.Body)
	var medicine entity.Medicine
	json.Unmarshal(bodyReq, &medicine)

	medicineRepository := repository.NewMedicineRepository(database.GetConnection())
	_, err := medicineRepository.InsertMedicine(ctx, int32(recipeIdInt), medicine)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"code":    http.StatusCreated,
		"message": "add medicine succesfull",
	})
}

func GetAllPharmacy(c *gin.Context) {
	ctx := context.Background()

	pharmacyRepository := repository.NewPharmacyRepository(database.GetConnection())
	result, err := pharmacyRepository.GetPharmacy(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(200, result)
}
