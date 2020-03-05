package controller

import (
	"database/sql"
	"e-portfolio-api/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TickerCtl struct {
	DB *sql.DB
}

func (t *TickerCtl) GetTickerList(c *gin.Context) {
	testTickerList := []model.Ticker{{Id: 0, Ticker: "MMM", Devidened: 5.88}}

	res, err := json.Marshal(testTickerList)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": string(res),
		"error":  nil,
	})
	return
}
