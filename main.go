package main

import (
	"crypto/tls"
	"github.com/rs/cors"
	"gopkg.in/ini.v1"
	"log"
	"net/http"
)

func main() {
	router := setupRouter()

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://goatest.bet"},
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

	certFile := "/ssl/cert.pem"
	keyFile := "/ssl/private_key.pem"

	var tlsConfig = &tls.Config{}
	tlsConfig.MinVersion = tls.VersionTLS12

	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatal("Erreur lors du chargement du certificat:", err)
	}
	tlsConfig.Certificates = []tls.Certificate{cert}

	server := &http.Server{
		Addr:      address,
		Handler:   router,
		TLSConfig: tlsConfig,
	}

	log.Println("Server started on " + address)
	err = server.ListenAndServeTLS(certFile, keyFile)
	if err != nil {
		log.Fatal("Erreur lors du démarrage du serveur:", err)
	}
}
