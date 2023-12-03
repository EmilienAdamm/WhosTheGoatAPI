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

func GetInstance() *MariaDBSingleton {
	if instance == nil {
		instance = &MariaDBSingleton{}
		instance.initPool()
	}
	return instance
}

func (m *MariaDBSingleton) initPool() {
	iniData, err := ini.Load("config.ini")
	if err != nil {
		log.Fatal("Erreur lors du chargement du fichier de configuration BDD:", err)
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

	db.SetMaxOpenConns(5) // Limite de connexions
	m.pool = db
}

func (m *MariaDBSingleton) Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := m.pool.Query(query, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
