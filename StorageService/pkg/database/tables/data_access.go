package tables

import (
	"StorageService/pkg/model"
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type DataAccess struct {
}

func (da *DataAccess) GetData(pool *pgxpool.Pool, chatID int64, service string) (model.ServiceData, error) {
	var tableData model.ServiceData
	rows, err := pool.Query(context.Background(), "SELECT service, login, password FROM data_services WHERE chat_id = $1 and service = $2", chatID, service)
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

func (da *DataAccess) AddData(pool *pgxpool.Pool, sd model.ServiceData) error {
	_, err := pool.Exec(context.Background(), "INSERT INTO data_services (chat_id, service, login, password) VALUES ($1, $2, $3, $4)", sd.ChatID, sd.Service, sd.Login, sd.Password)
	if err != nil {
		return err
	}
	return err
}

func (da *DataAccess) DelData(pool *pgxpool.Pool, sd model.ServiceData) error {
	_, err := pool.Exec(context.Background(), "DELETE FROM data_services WHERE chat_id = $1 AND service = $2", sd.ChatID, sd.Service)
	if err != nil {
		return err
	}
	return err
}
