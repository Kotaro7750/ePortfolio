package auth

import (
	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"os"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()

		credentials, err := google.CredentialsFromJSON(ctx, []byte(os.Getenv("FIREBASE_KEY")))
		if err != nil {
			log.Fatalf("error credentials from json: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			c.Abort()
			return
		}

		opt := option.WithCredentials(credentials)

		app, err := firebase.NewApp(ctx, nil, opt)
		if err != nil {
			log.Fatalf("error initializing app: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			c.Abort()
			return
		}

		firebaseClient, err := app.Auth(context.Background())
		_, err = app.Auth(context.Background())
		if err != nil {
			log.Fatalf("error getting Auth client: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			c.Abort()
			return
		}

		authorizationHeader := c.Request.Header.Get("Authorization")
		if authorizationHeader == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Authorization header doesn't exist",
			})
			c.Abort()
			return
		}

		if authorizationHeader[:7] != "Bearer " {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "not Bearer",
			})
			c.Abort()
			return
		}

		tokenString := authorizationHeader[7:]

		token, err := firebaseClient.VerifyIDToken(context.Background(), tokenString)
		if err != nil {
			log.Printf("error verifying ID token: %v\n", err)
			c.JSON(http.StatusForbidden, gin.H{
				"error": err,
			})
			c.Abort()
			return
		}

		log.Printf("Verified ID token: %v\n", token)
		c.Next()
	}
}
