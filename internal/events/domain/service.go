package domain

import "fmt"

type spotService struct{}

func NewSpotService() *spotService {
	return &spotService{}
}

func (s *spotService) GenerateSpots(event *Event, quantity int) error {

	for i := range quantity {
		spotName := fmt.Sprintf("%c%d", 'A'+i/10, i%10+1)
		spot, err := NewSpot(event, spotName)
		if err != nil {
			return err
		}

		event.Spots = append(event.Spots, *spot)
	}

	return nil
}
