package main

import (
	"database/sql"
	"gopkg.in/ini.v1"
	"log"
)

type MariaDBSingleton struct {
	pool *sql.DB
}

var instance *MariaDBSingleton

// GetInstance
//   - @name GetInstance
//   - @description Get the instance of the singleton
//   - @return *MariaDBSingleton
func GetInstance() *MariaDBSingleton {
	if instance == nil {
		instance = &MariaDBSingleton{}
		instance.initPool()
	}
	return instance
}

// initPool
//   - @name initPool
//   - @description Initialize the pool of connections. Called by GetInstance.
func (m *MariaDBSingleton) initPool() {
	iniData, err := ini.Load("config.ini")
	if err != nil {
		log.Fatal("Error while loading config file:", err)
	}
	section := iniData.Section("DATABASE")
	IP := section.Key("IP").String()
	PORT := section.Key("PORT").String()
	DB := section.Key("DATABASE").String()
	USER := section.Key("USERNAME").String()
	PWD := section.Key("PASSWORD").String()
	dbURL := USER + ":" + PWD + "@tcp(" + IP + ":" + PORT + ")/" + DB
	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(5)
	m.pool = db
}

// Query
//   - @name Query
//   - @description Execute a query on the database
//   - @param query string
//   - @param args []interface{}
//   - @return *sql.Rows
//   - @return error
func (m *MariaDBSingleton) Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := m.pool.Query(query, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
