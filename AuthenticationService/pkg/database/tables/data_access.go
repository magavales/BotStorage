package tables

import (
	"AuthenticationService/pkg/model"
	"github.com/jackc/pgx"
	"log"
)

type DataAccess struct {
}

func (da *DataAccess) GetPwd(pool *pgx.ConnPool, data string) (model.TableData, error) {
	var tableData model.TableData
	rows, err := pool.Query("SELECT password, salt FROM secret_pwd WHERE chat_id = $1", data)
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

func (da *DataAccess) SetPwd(pool *pgx.ConnPool, chatID, pwd, salt string) error {
	_, err := pool.Exec("INSERT INTO secret_pwd (password, salt, chat_id) VALUES ($1, $2, $3)", pwd, salt, chatID)
	if err != nil {
		return err
	}
	return err
}
