package eshop

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

var (
	priceURL   = "https://api.ec.nintendo.com/v1/price"
	priceLimit = 50
)

type Price struct{}

// paginated
func Prices(country string, nsuids []string) ([]*Price, error) {
	pages := [][]string{nsuids} // TODO: paginate nsuids every priceLimit elements

	results := make([]*Price, 0, len(nsuids))
	for _, page := range pages {
		prices, err := doPriceRequest(country, page)
		if err != nil {
			return nil, err
		}
		results = append(results, prices...)
	}

	return results, nil
}

// single page
func doPriceRequest(country string, nsuids []string) ([]*Price, error) {
	u := url.URL{
		Scheme: "https",
		Host:   "api.ec.nintendo.com",
		Path:   "v1/price",
		RawQuery: url.Values{
			"country": []string{country},
			"ids":     nsuids,
			"limit":   []string{strconv.Itoa(priceLimit)},
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

	b, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(b))
	return nil, nil
}
