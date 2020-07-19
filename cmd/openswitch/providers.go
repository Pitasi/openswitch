package main

import (
	"context"
	"encoding/json"
	"os"

	"github.com/sirupsen/logrus"
)

type Provider interface {
	ID() string
	Provide(context.Context) ([]Game, error)
}

var providers = []Provider{
	NewEuropeanEshop([]string{
		"IT",
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
	}),
}

func Fetch(ctx context.Context) {
	for _, p := range providers {
		log := logrus.WithField("provider", p.ID())

		// fetch games and prices for this provider
		games, err := p.Provide(ctx)
		if err != nil {
			log.WithError(err).Error("can't fetch games")
			continue
		}

		DatabaseStore(games)
	}
}

func DatabaseStore(games []Game) {
	f, _ := os.Create("/tmp/games.json")
	json.NewEncoder(f).Encode(games)
}
