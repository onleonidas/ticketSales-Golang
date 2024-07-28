package domain

import (
	"errors"
	"time"
)

type Rating string

const (
	RatingLivre Rating = "L"
	Rating10    Rating = "L10"
	Rating12    Rating = "L12"
	Rating14    Rating = "L14"
	Rating16    Rating = "L16"
	Rating18    Rating = "L18"
)

type Event struct {
	ID           string
	Name         string
	Location     string
	Organization string
	Rating       Rating
	Date         time.Time
	ImageURL     string
	Capacity     int
	Price        float64
	PartnerID    int
	Spots        []Spot
	Tickets      []Ticket
}

var (
	ErrEventNameRequired     = errors.New("event name is required")
	ErrEventDateInvalidError = errors.New("event date is invalid")
	ErrEventCapacityInvalid  = errors.New("event capacity is invalid")
	ErrEventPriceInvalid     = errors.New("event price is invalid")
)

func (e Event) Validate() error {
	if e.Name == "" {
		return ErrEventNameRequired
	}

	if e.Date.Before(time.Now()) {
		return ErrEventDateInvalidError
	}

	if e.Capacity <= 0 {
		return ErrEventCapacityInvalid
	}

	if e.Price <= 0 {
		return ErrEventPriceInvalid
	}

	return nil
}

func (e Event) AddSpot(s Spot) (*Spot, error) {
	spot, err := NewSpot(&e, s.Name)
	if err != nil {
		return nil, err
	}

	e.Spots = append(e.Spots, *spot)
	return spot, nil
}
