package domain

type EventRepository interface {
	FindEvents() ([]*Event, error)
	FindEventByID(id string) (*Event, error)
	FindSpotByID(id string) (*Spot, error)
	FindSpotsByEventID(eventID string) ([]*Spot, error)
	ReserveSpot(spot *Spot) error
}
