package main

import (
	"time"
)

type Game struct {
	// ProviderID is a unique identifier of the provider that fetched this
	// game.
	ProviderID string `db:"provider_id"`

	// ProviderGameID is a unique identifier of this game in the catalog of the
	// provider.
	ProviderGameID string `db:"provider_game_id"`

	Title       string  `db:"title"`
	Description string  `db:"description"`
	ImageURL    string  `db:"image_url"`
	Offers      []Offer `db:"offers"`
}

func (g *Game) AddOffer(o *Offer) {
	o.GameID = g.ProviderGameID
	o.GameProviderID = g.ProviderID
	g.Offers = append(g.Offers, *o)
}

type Offer struct {
	// ProviderID is a unique identifier of the provider that fetched this
	// price for a game.
	ProviderID string `json:"provider_id" db:"provider_id"`

	// RegularPrice is the price for buying this game.
	RegularPrice float32 `json:"regular_price" db:"regular_price"`

	DiscountPrice float32    `json:"discount_price,omitempty" db:"discount_price"`
	DiscountStart *time.Time `json:"discount_start,omitempty" db:"discount_start"`
	DiscountEnd   *time.Time `json:"discount_end,omitempty" db:"discount_end"`

	// BuyLink is a link to the page where this game can be bought from.
	BuyLink string `json:"buy_link,omitempty" db:"buy_link"`

	// GameID and GameProviderID refer to Game struct fields.
	GameID         string `db:"game_id"`
	GameProviderID string `db:"game_provider_id"`
}

func (o *Offer) Price() float32 {
	p := o.RegularPrice
	if !o.DiscountStart.IsZero() {
		p = o.DiscountPrice
	}
	return p
}
