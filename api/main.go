package main

import (
	"e-portfolio-api/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	//datasource := os.Getenv("DATABASE_DATASOURCE")
	//if datasource == "" {
	//	log.Fatal("Cannot get datasource for database.")
	//}

	//db, err := sql.Open("mysql", datasource)
	//if err != nil {
	//	log.Fatal("Cannot open database")
	//}
	//log.Printf("datasource is %s\n", datasource)
	//defer db.Close()

	//if err := db.Ping(); err != nil {
	//	log.Fatal(err.Error())
	//}

	router := gin.Default()

	//CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"Content-Type", "Authorization"}

	router.Use(cors.New(config))

	tickerCtl := controller.TickerCtl{DB: nil}

	router.GET("/ticker", tickerCtl.GetTickerList)

	Port := os.Getenv("PORT")
	router.Run(":" + Port)

}
