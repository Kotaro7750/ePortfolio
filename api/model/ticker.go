package model

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
)

type Ticker struct {
	Id             int      `json:"id"`
	Ticker         string   `json:"ticker"`
	Dividened      float64  `json:"dividened"`
	SectorId       int      `json:"sector_id"`
	Sector         string   `json:"sector"`
	AboutURL       string   `json:"about_url"`
	Color          string   `json:"color"`
	DividenedMonth []string `json:"dividened_month"`
}

func GetTickerList(db *sql.DB) ([]Ticker, error) {
	rows, err := db.Query("SELECT ticker.id, ticker.ticker,ticker.dividened,ticker.sector,sector.sector, ticker.about_url, ticker.color,ticker.dividened_month FROM ticker INNER JOIN sector ON ticker.sector = sector.id")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var tickers []Ticker

	for rows.Next() {
		var ticker Ticker
		var dividenedMonth string
		if err := rows.Scan(&ticker.Id, &ticker.Ticker, &ticker.Dividened, &ticker.SectorId, &ticker.Sector, &ticker.AboutURL, &ticker.Color, &dividenedMonth); err != nil {
			log.Printf("Query Error: %s", err.Error())
			return nil, err
		}
		ticker.DividenedMonth = strings.Split(dividenedMonth, ",")
		tickers = append(tickers, ticker)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Row Error: %s", err.Error())
		return nil, err
	}
	return tickers, nil
}

func AddTicker(db *sql.DB, t Ticker) error {
	ins, err := db.Prepare("INSERT INTO ticker (ticker,dividened,sector,about_url,color,dividened_month) VALUES ($1,$2,$3,$4,$5,$6)")
	if err != nil {
		log.Printf("Query Error: %s", err.Error())
		return err
	}

	_, err = ins.Exec(t.Ticker, t.Dividened, t.SectorId, t.AboutURL, t.Color, strings.Join(t.DividenedMonth, ","))
	if err != nil {
		log.Printf("Exec Error: %s", err.Error())
		return err
	}

	return nil
}

func UpdateTicker(db *sql.DB, t Ticker) error {
	ins, err := db.Prepare("UPDATE ticker SET ticker=$1,dividened=$2,sector=$3,about_url=$4,color=$5,dividened_month=$6 WHERE id=$7")
	if err != nil {
		log.Printf("Query Error: %s", err.Error())
		return err
	}

	_, err = ins.Exec(t.Ticker, t.Dividened, t.SectorId, t.AboutURL, t.Color, strings.Join(t.DividenedMonth, ","), t.Id)
	if err != nil {
		log.Printf("Exec Error: %s", err.Error())
		return err
	}

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
