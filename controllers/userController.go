package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	database "github.com/jain-vatsal/go-jwt/database"
	helpers "github.com/jain-vatsal/go-jwt/helpers"
	models "github.com/jain-vatsal/go-jwt/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(userPassword string, providedPassword string) (bool, string) {

	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg = fmt.Sprintf("email or password is incorrect")
		check = false
	}
	return check, msg
}

func SignUpUser() gin.HandlerFunc {

	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(user)

		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		countEmail, errEmail := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		defer cancel()

		if errEmail != nil {
			log.Panic(errEmail)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for the email"})
		}

		password := HashPassword(*user.Password)
		user.Password = &password
		countPhone, errPhone := userCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
		defer cancel()
		if errPhone != nil {
			log.Panic(errPhone)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for the phone number"})
		}

		if countEmail > 0 || countPhone > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "this email or phone number already exists"})
		}

		user.Created_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_ID = user.ID.Hex()
		token, refreshToken, _ := helpers.GenerateAllTokens(*user.Email, *user.First_Name, *user.Last_Name, *user.User_Type, *&user.User_ID)
		user.Token = &token
		user.Refresh_Token = &refreshToken

		resultInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)
		if insertErr != nil {
			msg := fmt.Sprintf("User item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		}

		defer cancel()
		c.JSON(http.StatusOK, resultInsertionNumber)
	}
}

func LoginUser() gin.HandlerFunc {

	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background())
		var user models.User
		var foundUser models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&user)
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "email or password is incorrect"})
			return
		}

		passwordIsValid, msg := VerifyPassword(*user.Password, *foundUser.Password)
		defer cancel()

		if passwordIsValid != true {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		if foundUser.Email != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
		}
		token, refreshToken, _ := helpers.GenerateAllTokens(*foundUser.Email, *foundUser.First_Name, *foundUser.Last_Name, *foundUser.User_Type, *&foundUser.User_ID)
		helpers.UpdateAllTokens(token, refreshToken, foundUser.User_ID)
		err = userCollection.FindOne(ctx, bson.M{"user_id": foundUser.User_ID}).Decode(&foundUser)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, foundUser)
	}

}

func GetAllUsers() {

}

func GetUser() gin.HandlerFunc {

	return func(c *gin.Context) {
		userID := c.Param("user_id")
		if err := helpers.MatchUserTypeToUID(c, userID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		err := userCollection.FindOne(ctx, bson.M{"user_id": userID}).Decode(&user)
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}
