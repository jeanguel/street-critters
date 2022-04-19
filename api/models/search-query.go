package models

import (
	"net/url"
	"strconv"
)

type SearchQuery struct {
	Page   int
	Amount int
}

func (sq SearchQuery) GetOffset() int {
	return (sq.Page - 1) * sq.Amount
}

// ParseParams
func (sq SearchQuery) ParseParams(query url.Values) (*SearchQuery, error) {
	pageStr := query.Get("page")
	if len(pageStr) > 0 {
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			return &sq, err
		}

		sq.Page = page
	} else {
		sq.Page = 1
	}

	amountStr := query.Get("amount")
	if len(amountStr) > 0 {
		amount, err := strconv.Atoi(amountStr)
		if err != nil {
			return &sq, err
		}

		sq.Amount = amount
	} else {
		sq.Amount = 10
	}

	return &sq, nil
}
