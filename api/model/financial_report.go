package model

import (
	"database/sql"
	"fmt"
	"log"
)

type FinancialReport struct {
	TickerId int    `json:"ticker_id"`
	Year     int    `json:"year"`
	Quarter  int    `json:"quarter"`
	URL      string `json:"url"`
}

func GetFinancialReportURL(db *sql.DB, ticker_id int, year int, quarter int) (string, error) {
	rows, err := db.Query(fmt.Sprintf("SELECT url FROM financial_report where ticker_id = %d AND year = %d AND quarter = %d", ticker_id, year, quarter))

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var url string

	for rows.Next() {
		if err := rows.Scan(&url); err != nil {
			log.Printf("Query Error: %s", err.Error())
			return "", err
		}
	}

	if err := rows.Err(); err != nil {
		log.Printf("Row Error: %s", err.Error())
		return "", err
	}
	return url, nil
}

func AddFinancialReport(db *sql.DB, f FinancialReport) error {
	ins, err := db.Prepare("INSERT INTO financial_report (ticker_id,year,quarter,url) VALUES ($1,$2,$3,$4)")
	if err != nil {
		log.Printf("Query Error: %s", err.Error())
		return err
	}

	_, err = ins.Exec(f.TickerId, f.Year, f.Quarter, f.URL)
	if err != nil {
		log.Printf("Exec Error: %s", err.Error())
		return err
	}

	return nil
}

func UpdateFinancialReport(db *sql.DB, f FinancialReport) error {
	ins, err := db.Prepare("UPDATE financial_report SET url=$1 WHERE ticker_id=$2 AND year = $3 AND quarter = $4")
	if err != nil {
		log.Printf("Query Error: %s", err.Error())
		return err
	}

	_, err = ins.Exec(f.URL, f.TickerId, f.Year, f.Quarter)
	if err != nil {
		log.Printf("Exec Error: %s", err.Error())
		return err
	}

	return nil
}
