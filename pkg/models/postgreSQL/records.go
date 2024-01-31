package postgreSQL

import (
	"database/sql"
	"good/goodservice/pkg/models"
	"log"
)

type GoodShopModel struct {
	DB *sql.DB
}

func (m *GoodShopModel) GET() (*[]models.Record, error) {
	query := `SELECT * FROM records`

	rows, err := m.DB.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	records := make([]models.Record, 0)

	for rows.Next() {
		r := models.Record{}
		if err := rows.Scan(&r.Artist, &r.Record_Name, &r.Record_Price); err != nil {
			log.Fatal(err)
		}
		records = append(records, r)
	}
	return &records, nil
}
