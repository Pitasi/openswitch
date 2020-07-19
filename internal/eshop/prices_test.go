package eshop

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPrice(t *testing.T) {
	res, err := Prices("IT", []string{"70010000000184"})
	assert := assert.New(t)
	assert.NoError(err)
	assert.Len(res, 1)
	assert.Equal(70010000000184, res[0].TitleID)
	log.Println(res[0].BuyLink)
}

func TestSplitIntoPages(t *testing.T) {
	tests := []struct {
		name     string
		list     []string
		expected [][]string
		pageSize int
	}{
		{
			name:     "nil list",
			list:     nil,
			expected: nil,
			pageSize: 10,
		},
		{
			name:     "empty list",
			list:     []string{},
			expected: nil,
			pageSize: 10,
		},
		{
			name: "three items three pages",
			list: []string{"1", "2", "3"},
			expected: [][]string{
				{"1"},
				{"2"},
				{"3"},
			},
			pageSize: 1,
		},
		{
			name: "three items two pages",
			list: []string{"1", "2", "3"},
			expected: [][]string{
				{"1", "2"},
				{"3"},
			},
			pageSize: 2,
		},
		{
			name: "page size larger than list",
			list: []string{"1", "2", "3"},
			expected: [][]string{
				{"1", "2", "3"},
			},
			pageSize: 100,
		},
		{
			name:     "zero page size",
			list:     []string{"1", "2", "3"},
			expected: nil,
			pageSize: 0,
		},
		{
			name:     "negative page size",
			list:     []string{"1", "2", "3"},
			expected: nil,
			pageSize: -1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := splitIntoPages(test.list, test.pageSize)
			assert.Equal(t, test.expected, result)
		})
	}
}
