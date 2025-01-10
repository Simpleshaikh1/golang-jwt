package controllers

import (
	"context"
	"fmt"
	"github.com/Simpleshaikh1/golang-jwt/database"
	"github.com/Simpleshaikh1/golang-jwt/helpers"
	helper "github.com/Simpleshaikh1/golang-jwt/helpers"
	"github.com/Simpleshaikh1/golang-jwt/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strconv"
	"time"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "users")
var validate = validator.New()

func HashPassword(password) {

}

func VerifyPassword() {

}

func Signup() {

}

func Login() {

}

func GetUsers() {

}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("user_id")

		if err := helper.MatchUserTypeToUid(c, userId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)

		defer cancel()

		var user models.User

		userCollection.FindOne()

	}
}
