// Package eshop implements methods to fetch data from Nintendo API.
//
// There are no known limits but every call should be considered "expensive"
// and used when necessary (i.e. use a cache!).
package eshop

import (
	"encoding/json"
	"net/http"
	"net/url"
)

// EuropeGames calls Nintendo API to fetch a list of all games for the European
// region.
func EuropeGames() ([]EuropeanGame, error) {
	u := url.URL{
		Scheme: "http",
		Host:   "search.nintendo-europe.com",
		Path:   "en/select", // en is one of the available locales
		RawQuery: url.Values{
			"rows":  []string{"9999"}, // max number of games returned
			"fq":    []string{"type:GAME AND system_type:nintendoswitch* AND product_code_txt:*"},
			"q":     []string{"*"},
			"sort":  []string{"sorting_title asc"},
			"start": []string{"0"},
			"wt":    []string{"json"},
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

	var apiRes struct {
		Response struct {
			Docs []EuropeanGame `json:"docs"`
		} `json:"response"`
	}

	err = json.NewDecoder(res.Body).Decode(&apiRes)
	if err != nil {
		return nil, err
	}

	return apiRes.Response.Docs, nil
}
