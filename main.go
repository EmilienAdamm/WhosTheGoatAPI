package main

import (
	"github.com/rs/cors"
	"gopkg.in/ini.v1"
	"log"
	"net/http"
)

func main() {
	router := setupRouter()

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowCredentials: true,
	})

	handler := corsHandler.Handler(router)

	startServer(handler)
}

func startServer(router http.Handler) {
	iniData, err := ini.Load("config.ini")
	if err != nil {
		log.Fatal("Erreur lors du chargement du fichier de configuration:", err)
	}

	section := iniData.Section("API")
	IP := section.Key("IP").String()
	PORT := section.Key("PORT").String()
	address := IP + ":" + PORT
	print("Server started on " + address + "\n")
	log.Fatal(http.ListenAndServe(address, router))
}
