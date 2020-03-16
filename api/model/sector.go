package model

import (
	"database/sql"
	"log"
)

type Sector struct {
	Id     int    `json:"id"`
	Sector string `json:"sector"`
}

func GetSectorList(db *sql.DB) ([]Sector, error) {
	rows, err := db.Query("SELECT id,sector FROM sector")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var sectors []Sector

	for rows.Next() {
		var sector Sector
		if err := rows.Scan(&sector.Id, &sector.Sector); err != nil {
			log.Printf("Query Error: %s", err.Error())
			return nil, err
		}
		sectors = append(sectors, sector)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Row Error: %s", err.Error())
		return nil, err
	}
	return sectors, nil
}
