package main

//"firebase.google.com/go/auth"
import (
	"database/sql"
	"e-portfolio-api/auth"
	"e-portfolio-api/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	datasource := os.Getenv("DATABASE_URL")
	if datasource == "" {
		log.Fatal("Cannot get datasource for database.")
	}

	db, err := sql.Open("postgres", datasource)
	if err != nil {
		log.Fatal("Cannot open database")
	}
	log.Printf("datasource is %s\n", datasource)
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err.Error())
	}

	router := gin.Default()

	//CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"Content-Type", "Authorization"}

	router.Use(cors.New(config))

	tickerCtl := controller.TickerCtl{DB: db}
	purchaseCtl := controller.PurchaseCtl{DB: db}

	router.GET("/ticker", auth.AuthMiddleWare(), tickerCtl.GetList)
	router.POST("/ticker", auth.AuthMiddleWare(), tickerCtl.Add)
	router.PUT("/ticker/:id", auth.AuthMiddleWare(), tickerCtl.Update)
	router.DELETE("/ticker/:id", auth.AuthMiddleWare(), tickerCtl.Delete)

	router.GET("/purchase", auth.AuthMiddleWare(), purchaseCtl.GetList)
	router.POST("/purchase", auth.AuthMiddleWare(), purchaseCtl.Add)
	router.PUT("/purchase/:id", auth.AuthMiddleWare(), purchaseCtl.Update)
	router.DELETE("/purchase/:id", auth.AuthMiddleWare(), purchaseCtl.Delete)

	Port := os.Getenv("PORT")
	router.Run(":" + Port)

}
