package domain

import (
	"errors"

	"github.com/google/uuid"
)

type SpotStatus string

const (
	SpotStatusAvailable SpotStatus = "available"
	SpotStatusReserved  SpotStatus = "reserved"
)

var (
	ErrSpotIDRequired       = errors.New("spot ID is required")
	ErrSpotNameRequired     = errors.New("spot name is required")
	ErrSpotEventIDMismatch  = errors.New("spot event ID mismatch")
	ErrSpotStatusRequired   = errors.New("spot status is required")
	ErrSpotTicketIDRequired = errors.New("spot ticket ID is required")
	ErrSpotNameInvalid      = errors.New("spot name is invalid, must start with a letter and end with a number")
	ErrSpotNotFound         = errors.New("spot not found")
)

type Spot struct {
	ID       string
	Name     string
	EventID  string
	Status   SpotStatus
	TicketID string
}

func NewSpot(event *Event, name string) (*Spot, error) {
	spot := &Spot{
		ID:      uuid.New().String(),
		Name:    name,
		EventID: event.ID,
		Status:  SpotStatusAvailable,
	}

	if err := spot.Validate(); err != nil {
		return nil, err
	}

	return spot, nil
}

func (s Spot) Validate() error {
	if s.ID == "" {
		return ErrSpotIDRequired
	}

	if s.Name == "" {
		return ErrSpotNameRequired
	}

	if s.Name[0] < 'A' || s.Name[0] > 'Z' {
		return ErrSpotNameInvalid
	}

	if s.Name[1] < '0' || s.Name[1] > '9' {
		return ErrSpotNameInvalid
	}

	return nil
}

func (s *Spot) Reserve(ticketID string) error {
	if s.Status != SpotStatusAvailable {
		return errors.New("spot is not available")
	}

	s.Status = SpotStatusReserved
	s.TicketID = ticketID

	return nil
}
