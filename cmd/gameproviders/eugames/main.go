package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/Pitasi/openswitch/internal/eshop"
)

var (
	list  *List
	mutex sync.Mutex
)

func main() {
	err := updateList()
	if err != nil {
		panic(err)
	}

	tick := time.NewTicker(1 * time.Hour)
	go func() {
		select {
		case <-tick.C:
			err := updateList()
			if err != nil {
				fmt.Printf("error during update: %v\n", err)
			}
		}
	}()

	http.HandleFunc("/", sendList)
	http.ListenAndServe(":8000", http.DefaultServeMux)
}

func updateList() error {
	log.Println("updating list")
	games, err := EuropeGames()
	if err != nil {
		return err
	}

	prices, err := fetchAllPrices(games)
	if err != nil {
		return err
	}
	mutex.Lock()
	defer mutex.Unlock()
	list = NewList(games, prices)
	log.Println("update completed succesfully")
	return nil
}

func fetchAllPrices(games []EuropeanGame) ([][]eshop.APIPrice, error) {
	return make([][]eshop.APIPrice, len(games)), nil
}

func sendList(w http.ResponseWriter, req *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	json.NewEncoder(w).Encode(list)
}
