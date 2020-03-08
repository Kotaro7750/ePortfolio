package controller

import (
	"database/sql"
	"e-portfolio-api/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ShareCtl struct {
	DB *sql.DB
}

func (s *ShareCtl) GetList(c *gin.Context) {
	shareList, err := model.GetShareList(s.DB)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	res, err := json.Marshal(shareList)

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
