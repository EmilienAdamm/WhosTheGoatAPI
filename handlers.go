package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net"
	"net/http"
	"time"
)

/**
 * @name setupRouter
 * @desc Routeer initialisation and endpoints definition
 * @return http.Handler
 */
func setupRouter() http.Handler {
	router := mux.NewRouter()
	// Endpoints de l'API
	router.HandleFunc("/ranPlayer", dealRanPlayer).Methods("GET")
	//router.HandleFunc("/updatePlayer", dealUpdatePlayer).Methods("POST")
	router.HandleFunc("/addScore", dealAddScore).Methods("POST")
	router.HandleFunc("/addAnalytics", dealAddAnalytics).Methods("POST")

	return router
}

func logRequest(request *http.Request) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	ip, _, _ := net.SplitHostPort(request.RemoteAddr)

	fmt.Printf("[%s] [%s] [%s] %s\n", currentTime, ip, request.Method, request.URL.Path)
}

func dealRanPlayer(response http.ResponseWriter, request *http.Request) {
	logRequest(request)

	DAO := NewDAO(GetInstance().pool)
	player, err := DAO.GetRanPlayer()
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(player)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	_, err = response.Write(jsonResponse)
	if err != nil {
		return
	}
}

func dealUpdatePlayer(response http.ResponseWriter, request *http.Request) {

}

func dealAddScore(response http.ResponseWriter, request *http.Request) {
	logRequest(request)

	var data ScoreData
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	DAO := NewDAO(GetInstance().pool)
	err = DAO.AddResult(data)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	_, err = response.Write([]byte("Score ajouté"))
	if err != nil {
		return
	}
}

func dealAddAnalytics(response http.ResponseWriter, request *http.Request) {
	logRequest(request)

	var data AnalyticsData
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	DAO := NewDAO(GetInstance().pool)
	err = DAO.AddAnalytics(data)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	_, err = response.Write([]byte("Analytics ajouté"))
	if err != nil {
		return
	}
}
