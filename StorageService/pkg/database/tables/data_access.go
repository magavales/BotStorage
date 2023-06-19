package tables

import (
	"StorageService/pkg/model"
	"github.com/jackc/pgx"
	"log"
)

type DataAccess struct {
}

func (da *DataAccess) GetData(pool *pgx.ConnPool, chatID string, service string) (model.ServiceData, error) {
	var tableData model.ServiceData
	rows, err := pool.Query("SELECT service, login, password FROM data_services WHERE chat_id = $1 and service = $2", chatID, service)
	if err != nil {
		log.Printf("The request was made incorrectly: %s\n", err)
	}

	if rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Println("error while iterating dataset")
		}
		tableData.ParseData(values)
		return tableData, nil
	}
	return tableData, pgx.ErrNoRows
}

func (da *DataAccess) AddData(pool *pgx.ConnPool, sd model.ServiceData) error {
	_, err := pool.Query("INSERT INTO data_services (chat_id, service, login, password) VALUES ($1, $2, $3, $4)", sd.ChatID, sd.Service, sd.Login, sd.Password)
	if err != nil {
		return err
	}
	return nil
}
