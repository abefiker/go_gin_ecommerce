package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/abefiker/go_gin_ecommerce/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func HashedPassword(password string) string {

}
func VerfyPassword(userPassword string, givenPassword string) (bool string) {

}
func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := Validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr})
			return
		}
		count, err := UserCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user already exist"})
			return
		}
		count, err := UserCollection.CountDocuments(ctx, bson.H{"phone": user.Phone})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "this phone already exist"})
		}
	}
}
func Login() gin.HandlerFunc {

}
func ProductViewAdmin() gin.HandlerFunc {

}
func SearchProduct() gin.HandlerFunc {

}
func SearchProductByQuery() gin.HandlerFunc {

}
