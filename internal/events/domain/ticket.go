package domain

import (
	"errors"

	"github.com/google/uuid"
)

type TicketKind string

const (
	TicketTypeFull TicketKind = "full"
	TicketTypeHalf TicketKind = "half"
)

var (
	ErrTicketPriceRequired = errors.New("ticket price not be minor or equal to 0")
	ErrInvalidTicketKind   = errors.New("invalid ticket kind")
)

type Ticket struct {
	ID         string
	EventID    string
	Spot       *Spot
	TicketKind TicketKind
	Price      float64
}

func NewTicket(event *Event, spot *Spot, ticketKind TicketKind) (*Ticket, error) {
	if !TicketValidate(ticketKind) {
		return nil, ErrInvalidTicketKind
	}

	ticket := &Ticket{
		ID:         uuid.New().String(),
		EventID:    event.ID,
		Spot:       spot,
		TicketKind: ticketKind,
		Price:      event.Price,
	}
	ticket.CalculatePrice()
	if err := ticket.Validate(); err != nil {
		return nil, err
	}
	return ticket, nil
}

func TicketValidate(ticket TicketKind) bool {
	return ticket == TicketTypeFull || ticket == TicketTypeHalf
}

func (t *Ticket) CalculatePrice() {
	if t.TicketKind == TicketTypeHalf {
		t.Price /= 2
	}
}

func (t *Ticket) Validate() error {
	if t.Price <= 0 {
		return ErrTicketPriceRequired
	}

	return nil
}
