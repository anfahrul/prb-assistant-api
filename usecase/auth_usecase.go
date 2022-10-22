package usecase

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/anfahrul/prb-assistant-api/database"
	"github.com/anfahrul/prb-assistant-api/entity"
	"github.com/anfahrul/prb-assistant-api/middlewares"
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

func Register(c *gin.Context) {
	ctx := context.Background()

	bodyReq, _ := ioutil.ReadAll(c.Request.Body)
	var input entity.User
	json.Unmarshal(bodyReq, &input)

	u := entity.User{
		Username:   input.Username,
		Password:   input.Password,
		Email:      input.Email,
		Role:       input.Role,
		IsLoggedIn: 0,
	}

	userRepository := repository.NewUserRepository(database.GetConnection())
	_, err := userRepository.Insert(ctx, u)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success"})
}

func Login(c *gin.Context) {
	middlewares.CORS()
	ctx := context.Background()

	bodyReq, _ := ioutil.ReadAll(c.Request.Body)
	var input entity.User
	json.Unmarshal(bodyReq, &input)

	userRepository := repository.NewUserRepository(database.GetConnection())
	token, err := userRepository.LoginCheck(ctx, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Kombinasi username & password salah",
		})
		return
	}

	tokenExpirationTime := time.Now().Add(time.Hour * 1)
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  tokenExpirationTime,
		Path:     "/",
		HttpOnly: false,
		Domain:   "api.prbassistant.test",
	})

	err = userRepository.UpdateLoginStatus(ctx, input, 1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "login success",
		"token":   token,
	})
}

func Logout(c *gin.Context) {
	ctx := context.Background()
	username := c.Request.Context().Value("username").(string)
	u := entity.User{
		Username: username,
	}

	userRepository := repository.NewUserRepository(database.GetConnection())
	err := userRepository.UpdateLoginStatus(ctx, u, 0)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token",
		Value:    "",
		MaxAge:   -1,
		Path:     "/",
		HttpOnly: false,
		SameSite: http.SameSiteStrictMode,
	})

	c.JSON(http.StatusOK, gin.H{"message": "logout success"})
}
