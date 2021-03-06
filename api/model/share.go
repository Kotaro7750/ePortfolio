package model

import (
	"database/sql"
	"log"
	"strings"
)

type ShareForTicker struct {
	Id             int      `json:"id"`
	Ticker         string   `json:"ticker"`
	Sector         string   `json:"sector"`
	Color          string   `json:"color"`
	Amount         int      `json:"amount"`
	Dividened      float64  `json:"dividened"`
	TotalCost      float64  `json:"total_cost"`
	MeanCost       float64  `json:"mean_cost"`
	DividenedMonth []string `json:"dividened_month"`
}

type ShareForSector struct {
	Id        int     `json:"id"`
	Sector    string  `json:"sector"`
	Color     string  `json:"color"`
	Amount    int     `json:"amount"`
	Dividened float64 `json:"dividened"`
	TotalCost float64 `json:"total_cost"`
	MeanCost  float64 `json:"mean_cost"`
}

func GetShareForTicker(db *sql.DB) ([]ShareForTicker, error) {
	rows, err := db.Query(
		`SELECT 

      ticker.id,
      ticker.ticker,
      sector.sector,
      ticker.color,
      SUM(purchase_history.share * ticker.dividened),
      ticker.dividened_month,
      SUM(purchase_history.share),
      SUM(purchase_history.cost) 

     FROM purchase_history INNER JOIN ticker ON purchase_history.ticker_id = ticker.id 
     INNER JOIN sector ON ticker.sector = sector.id
     GROUP BY ticker.id,sector.sector
    `)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var shares []ShareForTicker

	for rows.Next() {
		var share ShareForTicker
		var dividenedMonth string

		if err := rows.Scan(&share.Id, &share.Ticker, &share.Sector, &share.Color, &share.Dividened, &dividenedMonth, &share.Amount, &share.TotalCost); err != nil {
			log.Printf("Query Error: %s", err.Error())
			return nil, err
		}

		share.DividenedMonth = strings.Split(dividenedMonth, ",")
		share.MeanCost = share.TotalCost / float64(share.Amount)
		shares = append(shares, share)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Row Error: %s", err.Error())
		return nil, err
	}

	return shares, nil
}

func GetShareForSector(db *sql.DB) ([]ShareForSector, error) {
	rows, err := db.Query(
		`SELECT 

      sector.id,
      sector.sector,
      sector.color,
      SUM(purchase_history.share * ticker.dividened),
      SUM(purchase_history.share),
      SUM(purchase_history.cost) 

     FROM purchase_history INNER JOIN ticker ON purchase_history.ticker_id = ticker.id 
     INNER JOIN sector ON ticker.sector = sector.id
     GROUP BY sector.id
    `)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var shares []ShareForSector

	for rows.Next() {
		var share ShareForSector

		if err := rows.Scan(&share.Id, &share.Sector, &share.Color, &share.Dividened, &share.Amount, &share.TotalCost); err != nil {
			log.Printf("Query Error: %s", err.Error())
			return nil, err
		}
		share.MeanCost = share.TotalCost / float64(share.Amount)
		shares = append(shares, share)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Row Error: %s", err.Error())
		return nil, err
	}

	return shares, nil
}
