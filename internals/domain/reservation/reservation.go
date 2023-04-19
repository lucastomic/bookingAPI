package reservation

import (
	"time"

	"github.com/lucastomic/naturalYSalvajeRent/internals/domain/user"
)

type Reservation struct {
	user     user.User
	firstDay time.Time
	lastDay  time.Time
}
