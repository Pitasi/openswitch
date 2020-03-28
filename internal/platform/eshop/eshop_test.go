package eshop

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEuropeGames(t *testing.T) {
	games, err := EuropeGames()
	assert.NoError(t, err)
	assert.Greater(t, len(games), 1000) // there are at least 1000 games
}
