package main

import (
	"time"

	"github.com/Pitasi/openswitch/internal/eshop"
)

type List struct {
	FetchedAt time.Time
	Items     []*OutputGame
}

func NewList(games []EuropeanGame, prices [][]eshop.Price) *List {
	outGames := make([]*OutputGame, len(games))
	for i, g := range games {
		outGames[i] = NewOutputGame(g, prices[i])
	}
	l := &List{
		FetchedAt: time.Now(),
		Items:     outGames,
	}
	return l
}

type OutputGame struct {
	Title  string
	Prices []struct {
		Provider string
		URL      string
		Price    int
	}
}

func NewOutputGame(g EuropeanGame, p []eshop.Price) *OutputGame {
	return &OutputGame{
		Title: g.Title,
	}
}
