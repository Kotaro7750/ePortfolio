package model

import (
	"database/sql"
	"log"
)

type Share struct {
	Id        int     `json:"id"`
	Ticker    string  `json:"ticker"`
	Sector    string  `json:"sector"`
	Color     string  `json:"color"`
	Amount    int     `json:"amount"`
	Dividened float64 `json:"dividened"`
	TotalCost float64 `json:"total_cost"`
	MeanCost  float64 `json:"mean_cost"`
}

func GetShareList(db *sql.DB) ([]Share, error) {
	rows, err := db.Query(
		`SELECT 

      ticker.id,
      ticker.ticker,
      sector.sector,
      ticker.color,
      SUM(purchase_history.share * ticker.dividened),
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

	var shares []Share

	for rows.Next() {
		var share Share

		if err := rows.Scan(&share.Id, &share.Ticker, &share.Sector, &share.Color, &share.Dividened, &share.Amount, &share.TotalCost); err != nil {
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
