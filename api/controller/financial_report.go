package controller

import (
	"bytes"
	"database/sql"
	"e-portfolio-api/model"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
)

type FinancialReportCtl struct {
	DB *sql.DB
}

func (f *FinancialReportCtl) GetMD(c *gin.Context) {
	url := c.Query("url")

	response, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer response.Body.Close()

	buf := new(bytes.Buffer)
	io.Copy(buf, response.Body)

	c.Header("Content-Description", "File Transfer")
	c.Data(http.StatusOK, "text/plain", buf.Bytes())
	return
}

func (f *FinancialReportCtl) GetURL(c *gin.Context) {
	tickerId, err := strconv.Atoi(c.Param("ticker_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	year, err := strconv.Atoi(c.Param("year"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	quarter, err := strconv.Atoi(c.Param("quarter"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	url, err := model.GetFinancialReportURL(f.DB, tickerId, year, quarter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if url == "" {
		c.JSON(http.StatusOK, gin.H{
			"url":        "",
			"is_url_set": false,
			"error":      nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"url":        url,
		"is_url_set": true,
		"error":      nil,
	})
	return
}

func (f *FinancialReportCtl) Add(c *gin.Context) {
	var financialReport model.FinancialReport

	if err := c.ShouldBindJSON(&financialReport); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if financialReport.URL == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "url must not be empty",
		})
		return
	}

	err := model.AddFinancialReport(f.DB, financialReport)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": financialReport,
		"error":  nil,
	})
	return
}

func (f *FinancialReportCtl) Update(c *gin.Context) {
	var financialReport model.FinancialReport

	if err := c.ShouldBindJSON(&financialReport); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := model.UpdateFinancialReport(f.DB, financialReport)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": financialReport,
		"error":  nil,
	})
	return
}
