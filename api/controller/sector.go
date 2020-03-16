package controller

import (
	"database/sql"
	"e-portfolio-api/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SectorCtl struct {
	DB *sql.DB
}

func (s *SectorCtl) GetList(c *gin.Context) {
	sectorList, err := model.GetSectorList(s.DB)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	res, err := json.Marshal(sectorList)

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
