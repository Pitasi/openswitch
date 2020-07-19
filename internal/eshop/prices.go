package eshop

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var maxPageSize = 50

type APIPrice struct {
	// SalesStatus says if the game is available to buy:
	// Currently known values:
	// "onsale":     available to buy
	// "unreleased": not available to buy, there are no prices yet
	// "not_found":  the game is not available in the requested region
	SalesStatus string `json:"sales_status"`

	// TitleID is the NSUID of the game
	TitleID int `json:"title_id"`

	// RegularPrice is the price for this game.
	// It can be empty if the game hasn't been released yet, use IsOnSale() to
	// check.
	RegularPrice APIRegularPrice `json:"regular_price"`

	// DiscountPrice is the discounted price for this game.
	// It is empty if there are no ongoing discounts, use IsDiscounted() to
	// check.
	DiscountPrice APIDiscountPrice `json:"discount_price"`

	// BuyLink is the full address to buy the game at the price listed.
	//
	// Note: this field is not actual part of the API response, it's generated
	// by this package for convenience.
	BuyLink string `json:"buy_link"`
}

// IsOnSale returns true if the game is available to buy and has a
// RegularPrice.
func (p *APIPrice) IsOnSale() bool {
	return p.SalesStatus == "onsale"
}

// IsDiscounted returns true if there's a discount for the game, and has a
// DiscountPrice.
func (p *APIPrice) IsDiscounted() bool {
	return p.DiscountPrice.StartDatetime.IsZero()
}

type APIRegularPrice struct {
	// Amount is the price combined with the currency symbol.
	// E.g. "39.99 €"
	Amount string `json:"amount"`

	// Currency is the three-characters name of the currency being used.
	// E.g. "EUR"
	Currency string `json:"currency"`

	// RawValue is a string representation of the price.
	// E.g. "39.99"
	RawValue string `json:"raw_value"`
}

type APIDiscountPrice struct {
	APIRegularPrice
	StartDatetime time.Time
	EndDatetime   time.Time
}

type APIPriceResponse struct {
	Country      string
	Personalized bool
	Prices       []*APIPrice

	// Error is set when something "bad" happened.
	Error *APIError `json:"error"`
}

type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("API Error: code=%s, message=%s", e.Code, e.Message)
}

// Prices executes several calls to Nintendo API for fetching prices related
// to requested NSUIDs. Each call to the API can retrieve up to `maxPageSize`
// results.
func Prices(country string, nsuids []string) ([]*APIPrice, error) {
	country = strings.ToUpper(country)
	pages := splitIntoPages(nsuids, maxPageSize)

	results := make([]*APIPrice, 0, len(nsuids))
	for _, page := range pages {
		prices, err := doPriceRequest(country, page)
		if err != nil {
			return nil, err
		}
		results = append(results, prices...)
	}

	fillLinkURLs(country, results)

	return results, nil
}

// doPriceRequest executes a single request to Nintendo API for fetching
// prices. Length of nsuids must not exceed `maxPageSize`.
func doPriceRequest(country string, nsuids []string) ([]*APIPrice, error) {
	if len(nsuids) > maxPageSize {
		return nil, fmt.Errorf("requested %d prices but maximum is %d", len(nsuids), maxPageSize)
	}

	u := url.URL{
		Scheme: "https",
		Host:   "api.ec.nintendo.com",
		Path:   "v1/price",
		RawQuery: url.Values{
			"country": []string{country},
			"ids":     []string{strings.Join(nsuids, ",")},
			"limit":   []string{strconv.Itoa(maxPageSize)},
			"lang":    []string{"en"},
		}.Encode(),
	}

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	price := new(APIPriceResponse)
	err = json.NewDecoder(res.Body).Decode(price)
	if err != nil {
		return nil, err
	}
	if price.Error != nil {
		return nil, price.Error
	}

	return price.Prices, nil
}

func splitIntoPages(l []string, pageSize int) [][]string {
	if len(l) == 0 || pageSize <= 0 {
		return nil
	}

	if pageSize > len(l) {
		return [][]string{l}
	}

	nPages := int(math.Ceil(float64(len(l)) / float64(pageSize)))
	pages := make([][]string, nPages)

	for i := 0; i < nPages; i++ {
		bound := (i + 1) * pageSize
		if bound > len(l) {
			bound = len(l)
		}
		pages[i] = l[i*pageSize : bound]
	}

	return pages
}

func fillLinkURLs(country string, res []*APIPrice) {
	format := "https://ec.nintendo.com/%s/en/titles/%d"
	for _, p := range res {
		p.BuyLink = fmt.Sprintf(format, country, p.TitleID)
	}
}
