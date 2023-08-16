package reservationrequest

type ReservationRequest struct {
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	FirstDay   string `json:"firstDay"`
	LastDay    string `json:"lastDay"`
	BoatId     int    `json:"boatId"`
	Passengers int    `json:"passengers"`
	IsOpen     bool   `json:"isOpen"`
}
