package main

import (
	"time"
)

type Game struct {
	// ProviderID is a unique identifier of the provider that fetched this
	// game.
	ProviderID string

	// ProviderGameID is a unique identifier of this game in the catalog of the
	// provider.
	ProviderGameID string

	Title       string
	Description string
	ImageURL    string
	Offers      []Offer
}

func (g *Game) AddOffer(o Offer) {
	g.Offers = append(g.Offers, o)
}

type Offer struct {
	// ProviderID is a unique identifier of the provider that fetched this
	// price for a game.
	ProviderID string `json:"provider_id"`

	// RegularPrice is the price for buying this game.
	RegularPrice float32 `json:"regular_price"`

	DiscountPrice float32    `json:"discount_price,omitempty"`
	DiscountStart *time.Time `json:"discount_start,omitempty"`
	DiscountEnd   *time.Time `json:"discount_end,omitempty"`

	// BuyLink is a link to the page where this game can be bought from.
	BuyLink string `json:"buy_link,omitempty"`
}

func (o *Offer) Price() float32 {
	p := o.RegularPrice
	if !o.DiscountStart.IsZero() {
		p = o.DiscountPrice
	}
	return p
}
