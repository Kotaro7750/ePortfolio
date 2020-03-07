package controller

import (
	"database/sql"
	"e-portfolio-api/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TickerCtl struct {
	DB *sql.DB
}

func (t *TickerCtl) GetList(c *gin.Context) {
	tickerList, err := model.GetTickerList(t.DB)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	res, err := json.Marshal(tickerList)

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

func (t *TickerCtl) Add(c *gin.Context) {
	var ticker model.Ticker
	if err := c.ShouldBindJSON(&ticker); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := model.AddTicker(t.DB, ticker)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": ticker,
		"error":  nil,
	})
	return
}

func (t *TickerCtl) Update(c *gin.Context) {
	var ticker model.Ticker

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := c.ShouldBindJSON(&ticker); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if ticker.Id != id {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id is contradicted to URL",
		})
		return
	}

	err = model.UpdateTicker(t.DB, ticker)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": ticker,
		"error":  nil,
	})
	return

}

func (t *TickerCtl) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = model.DeleteTicker(t.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
}
