package database

import (
	"AuthenticationService/pkg/database/tables"
	"github.com/jackc/pgx"
	"log"
)

type PostgresDB struct {
	Pool   *pgx.ConnPool
	Access tables.DataAccess
}

func (db *PostgresDB) Connect() {
	config := pgx.ConnConfig{
		Host:     "localhost",
		Port:     5432,
		Database: "postgres",
		User:     "postgres",
		Password: "1703",
	}
	poolConn := pgx.ConnPoolConfig{
		ConnConfig:     config,
		MaxConnections: 5,
		AfterConnect:   nil,
		AcquireTimeout: 0,
	}

	var err error
	db.Pool, err = pgx.NewConnPool(poolConn)
	if err != nil {
		log.Printf("I can't connect to database: %s\n", err)
	}
}

func (db *PostgresDB) StatConn() {
	if db.Pool.Stat().MaxConnections == 4 {
		db.Pool.Close()
	}
}
