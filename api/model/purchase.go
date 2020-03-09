package model

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type Purchase struct {
	Id       int     `json:"id"`
	TickerId int     `json:"ticker_id"`
	Ticker   string  `json:"ticker"`
	Share    int     `json:"share"`
	Cost     float64 `json:"cost"`
	Date     string  `json:"date"`
}

func GetPurchaseList(db *sql.DB) ([]Purchase, error) {
	rows, err := db.Query(
		"SELECT purchase_history.id, ticker.id,ticker.ticker, purchase_history.date, purchase_history.share, purchase_history.cost FROM purchase_history INNER JOIN ticker ON purchase_history.ticker_id = ticker.id")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var purchases []Purchase

	for rows.Next() {
		var purchase Purchase
		if err := rows.Scan(&purchase.Id, &purchase.TickerId, &purchase.Ticker, &purchase.Date, &purchase.Share, &purchase.Cost); err != nil {
			log.Printf("Query Error: %s", err.Error())
			return nil, err
		}
		purchases = append(purchases, purchase)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Row Error: %s", err.Error())
		return nil, err
	}
	return purchases, nil
}

func AddPurchase(db *sql.DB, p Purchase) error {
	ins, err := db.Prepare("INSERT INTO purchase_history (ticker_id,date,share,cost) VALUES ($1,$2,$3,$4)")
	if err != nil {
		log.Printf("Query Error: %s", err.Error())
		return err
	}

	_, err = ins.Exec(p.TickerId, p.Date, p.Share, p.Cost)
	if err != nil {
		log.Printf("Exec Error: %s", err.Error())
		return err
	}

	return nil
}

func UpdatePurchase(db *sql.DB, p Purchase) error {
	ins, err := db.Prepare("UPDATE purchase_history SET ticker_id=$1,date=$2,share=$3,cost=$4 WHERE id=$5")
	if err != nil {
		log.Printf("Query Error: %s", err.Error())
		return err
	}

	_, err = ins.Exec(p.TickerId, p.Date, p.Share, p.Cost, p.Id)
	if err != nil {
		log.Printf("Exec Error: %s", err.Error())
		return err
	}

	return nil

}

func DeletePurchase(db *sql.DB, id int) error {
	res, err := db.Exec(fmt.Sprintf("DELETE FROM purchase_history WHERE id = '%d'", id))

	if err != nil {
		return err
	}

	if affectedRow, _ := res.RowsAffected(); affectedRow == 0 {
		return errors.New("no rows affected")
	}
	return nil
}
