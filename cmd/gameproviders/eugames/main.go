package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Pitasi/openswitch/internal/eshop"
)

func main() {
	start := time.Now()

	games, err := EuropeGames()
	if err != nil {
		panic(fmt.Errorf("error fetching games list: %w", err))
	}

	nsuids := make([]string, 0, len(games))
	for _, g := range games {
		nsuid, err := g.NSUID()
		if err != nil {
			log.Printf("no nsuid for %s, skipping price fetch\n", g.Title)
		}
		nsuids = append(nsuids, nsuid)
	}

	for _, country := range countries {
		prices, err := eshop.Prices(country, nsuids)
		if err != nil {
			panic(fmt.Errorf("error fetching prices: %w", err))
		}
		log.Println(country, len(prices), "games")
	}

	end := time.Now()
	duration := end.Sub(start)

	log.Printf("Done: fetched %d games in %s.", len(games), duration)
}

var countries = []string{
	"FI",
	"IE",
	"NO",
	"SE",
	"CY",
	"GR",
	"PT",
	"SI",
	"CZ",
	"HU",
	"RU",
	"SK",
	"FR",
	"DE",
	"CH",
	"AU",
	"DK",
	"EE",
	"LV",
	"LT",
	"GB",
	"HR",
	"IT",
	"MT",
	"ES",
	"BG",
	"PL",
	"RO",
	"AT",
	"BE",
	"LU",
	"NL",
	"NZ",
	"ZA",
}
