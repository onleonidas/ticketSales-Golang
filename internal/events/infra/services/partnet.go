package services

type ReservationResponse struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	EventID    string `json:"event_id"`
	Status     string `json:"status"`
	TicketKind string `json:"ticket_kind"`
	Spot       string `json:"spot"`
}

type ReservationRequest struct {
	Spots      []string `json:"spots"`
	EventID    string   `json:"event_id"`
	TicketKind string   `json:"ticket_type"`
	Email      string   `json:"email"`
	CardHash   string   `json:"card_hash"`
}

type Partner interface {
	MakeReservation(req *ReservationRequest) ([]ReservationResponse, error)
}
