package reservationrequest

type ReservationRequest struct {
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	FirstDay    string `json:"firstDay"`
	LastDay     string `json:"lastDay"`
	BoatId      int    `json:"boatId"`
	StateRoomId int    `json:"stateRoomId"`
}
