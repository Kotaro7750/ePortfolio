package model

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type Memo struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	AboutURL string `json:"about_url"`
}

func GetMemoList(db *sql.DB) ([]Memo, error) {
	rows, err := db.Query("SELECT id,title,url FROM memo")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var memos []Memo

	for rows.Next() {
		var memo Memo
		if err := rows.Scan(&memo.Id, &memo.Title, &memo.AboutURL); err != nil {
			log.Printf("Query Error: %s", err.Error())
			return nil, err
		}
		memos = append(memos, memo)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Row Error: %s", err.Error())
		return nil, err
	}
	return memos, nil
}

func AddMemo(db *sql.DB, m Memo) error {
	ins, err := db.Prepare("INSERT INTO memo (title,url) VALUES ($1,$2)")
	if err != nil {
		log.Printf("Query Error: %s", err.Error())
		return err
	}

	_, err = ins.Exec(m.Title, m.AboutURL)
	if err != nil {
		log.Printf("Exec Error: %s", err.Error())
		return err
	}

	return nil
}

func UpdateMemo(db *sql.DB, m Memo) error {
	ins, err := db.Prepare("UPDATE memo SET title=$1,url=$2 WHERE id=$3")
	if err != nil {
		log.Printf("Query Error: %s", err.Error())
		return err
	}

	_, err = ins.Exec(m.Title, m.AboutURL, m.Id)
	if err != nil {
		log.Printf("Exec Error: %s", err.Error())
		return err
	}

	return nil

}

func DeleteMemo(db *sql.DB, id int) error {
	res, err := db.Exec(fmt.Sprintf("DELETE FROM memo WHERE id = '%d'", id))

	if err != nil {
		return err
	}

	if affectedRow, _ := res.RowsAffected(); affectedRow == 0 {
		return errors.New("no rows affected")
	}
	return nil
}
