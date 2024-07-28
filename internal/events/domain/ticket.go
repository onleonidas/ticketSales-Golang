package domain

import "errors"

type TicketType string

const (
	TicketTypeFull TicketType = "full"
	TicketTypeHalf TicketType = "half"
)

var (
	ErrTicketPriceRequired = errors.New("ticket price not be minor or equal to 0")
)

type Ticket struct {
	ID         string
	EventID    string
	Spot       *Spot
	TicketType TicketType
	Price      float64
}

func ticketValidate(ticket TicketType) bool {
	return ticket == TicketTypeFull || ticket == TicketTypeHalf
}

func (t *Ticket) calculatePrice() {
	if t.TicketType == TicketTypeHalf {
		t.Price /= 2
	}
}

func (t *Ticket) Validate() error {
	if t.Price <= 0 {
		return ErrTicketPriceRequired
	}

	return nil
}
