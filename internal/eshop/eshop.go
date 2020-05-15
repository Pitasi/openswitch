// Package eshop implements methods to fetch data from Nintendo API.
//
// There are no known limits but every call should be considered "expensive"
// and used when necessary (i.e. use a cache!).
package eshop

import (
	"fmt"
)

func NSUID(game interface{}) (string, error) {
	switch g := game.(type) {
	case AmericanGame:
		return g.Nsuid, nil

	case AsianGame:
		// TODO: use a regex to extract it from g.LinkURL
		return "", fmt.Errorf("not implemented")
	}

	panic(fmt.Errorf("invalid game type"))
}
