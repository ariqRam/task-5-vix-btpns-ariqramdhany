package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/ramdhanyA/task-5-vix-btpns/database"
	helper "github.com/ramdhanyA/task-5-vix-btpns/helpers"
	"github.com/ramdhanyA/task-5-vix-btpns/models"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strconv"
	"time"
)

var users []models.User
var db = database.DBinstance()
var query = db.Find(&users)

func HashPassword() {}

func VerifyPassword() {}

func Signup() {}

func Login() {}

func GetUsers() {}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("user_id")

		if err := helper.MatchUserTypeToUid(c, userId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var user models.User
		result := db.First(&user, "user_id = ?", userId)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}
