package mockups

import "github.com/lucastomic/naturalYSalvajeRent/internals/domain"

type BoatMockUp struct{}

func (mock BoatMockUp) Save(boat domain.Boat) error                   { return nil }
func (mock BoatMockUp) Remove(boat domain.Boat) error                 { return nil }
func (mock BoatMockUp) FindById(ids ...int) (domain.Boat, error)      { return domain.Boat{}, nil }
func (mock BoatMockUp) FindAll() ([]domain.Boat, error)               { return []domain.Boat{}, nil }
func (mock BoatMockUp) FindByUser(boat string) ([]domain.Boat, error) { return []domain.Boat{}, nil }
