// Create and manage database.
package db

import (
	"database/sql"
	"fmt"
	"strings"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteDatastore struct {
	sqlDB *sql.DB
	Videos *VideoDB
}

func CreateDB(dbPath string, pin string) (*SQLiteDatastore, error) {
	//dbPath := path.Join(dbDir, "mainnet.db")
	conn, err :=sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	conn.SetMaxIdleConns(2)
	conn.SetMaxOpenConns(4)
	if pin != "" {
		p := "pragma key='" + strings.Replace(pin, "'", "''", -1) + "';"
		if _, err := conn.Exec(p); err != nil {
			return nil, err
		}
	}
	lock := new(sync.Mutex)
	return &SQLiteDatastore{sqlDB: conn, Videos: NewVideoDB(conn, lock)}, nil
}

func (s *SQLiteDatastore)InitDatabaseTables(pin string) error {
	var sqlStmt string
	if pin != "" {
		sqlStmt = "pragma key = '" + strings.Replace(pin, "'", "''", -1) + "';"
	}
	sqlStmt += `
	create table videos (ind integer primary key not null, videoName text not null, blockNum integer not null, poster text, videoLength integer, created integer not null, updated integer not null);
	create table blocks (ind integer primary key not null, id text not null, belongTo integer not null);
	create index block_belong on blocks(belongTo);
    `
	if _, err := s.sqlDB.Exec(sqlStmt); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}