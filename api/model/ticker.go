package model

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type Ticker struct {
	Id        int     `json:"id"`
	Ticker    string  `json:"ticker"`
	Dividened float64 `json:"dividened"`
}

func GetTickerList(db *sql.DB) ([]Ticker, error) {
	rows, err := db.Query("SELECT id,ticker,dividened FROM ticker")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var tickers []Ticker

	for rows.Next() {
		var ticker Ticker
		if err := rows.Scan(&ticker.Id, &ticker.Ticker, &ticker.Dividened); err != nil {
			log.Printf("Query Error: %s", err.Error())
			return nil, err
		}
		tickers = append(tickers, ticker)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Row Error: %s", err.Error())
		return nil, err
	}
	return tickers, nil
}

func AddTicker(db *sql.DB, t Ticker) error {
	ins, err := db.Prepare("INSERT INTO ticker (ticker,dividened) VALUES ($1,$2)")
	if err != nil {
		log.Printf("Query Error: %s", err.Error())
		return err
	}
	ins.Exec(t.Ticker, t.Dividened)
	return nil
}

func DeleteTicker(db *sql.DB, id int) error {
	res, err := db.Exec(fmt.Sprintf("DELETE FROM ticker WHERE id = '%d'", id))

	if err != nil {
		return err
	}

	if affectedRow, _ := res.RowsAffected(); affectedRow == 0 {
		return errors.New("no rows affected")
	}
	return nil
}
