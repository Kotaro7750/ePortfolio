package model

import (
	"database/sql"
	"log"
)

type Share struct {
	Id        int     `json:"id"`
	Ticker    string  `json:"ticker"`
	Amount    int     `json:"amount"`
	Dividened float64 `json:"dividened"`
	TotalCost float64 `json:"total_cost"`
	MeanCost  float64 `json:"mean_cost"`
}

func GetShareList(db *sql.DB) ([]Share, error) {
	rows, err := db.Query(
		"SELECT ticker.id,ticker.ticker,ticker.dividened,purchase_history.share, purchase_history.cost FROM purchase_history INNER JOIN ticker ON purchase_history.ticker_id = ticker.id")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var shares []Share
	tmp := make(map[int]struct {
		ticker    string
		amount    int
		dividened float64
		totalCost float64
	})

	for rows.Next() {
		var id int
		var ticker string
		var amount int
		var dividened float64
		var totalCost float64

		if err := rows.Scan(&id, &ticker, &dividened, &amount, &totalCost); err != nil {
			log.Printf("Query Error: %s", err.Error())
			return nil, err
		}
		amount += tmp[id].amount
		totalCost += tmp[id].totalCost

		tmp[id] = struct {
			ticker    string
			amount    int
			dividened float64
			totalCost float64
		}{
			ticker:    ticker,
			amount:    amount,
			dividened: dividened,
			totalCost: totalCost,
		}
	}

	if err := rows.Err(); err != nil {
		log.Printf("Row Error: %s", err.Error())
		return nil, err
	}

	for id, obj := range tmp {
		totalDividened := obj.dividened * float64(obj.amount)
		meanCost := obj.totalCost / float64(obj.amount)

		shares = append(shares, Share{Id: id, Ticker: obj.ticker, Amount: obj.amount, Dividened: totalDividened, TotalCost: obj.totalCost, MeanCost: meanCost})
	}
	return shares, nil
}
