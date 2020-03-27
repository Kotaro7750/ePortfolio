package controller

import (
	"database/sql"
	"e-portfolio-api/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type MemoCtl struct {
	DB *sql.DB
}

func (m *MemoCtl) GetList(c *gin.Context) {
	memoList, err := model.GetMemoList(m.DB)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	res, err := json.Marshal(memoList)

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

func (m *MemoCtl) Add(c *gin.Context) {
	var memo model.Memo
	if err := c.ShouldBindJSON(&memo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := model.AddMemo(m.DB, memo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": memo,
		"error":  nil,
	})
	return
}

func (m *MemoCtl) Update(c *gin.Context) {
	var memo model.Memo

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := c.ShouldBindJSON(&memo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if memo.Id != id {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id is contradicted to URL",
		})
		return
	}

	err = model.UpdateMemo(m.DB, memo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": memo,
		"error":  nil,
	})
	return

}

func (m *MemoCtl) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = model.DeleteMemo(m.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
}
