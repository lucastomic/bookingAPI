package viewport

import (
	"github.com/gin-gonic/gin"
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
	"github.com/lucastomic/naturalYSalvajeRent/internals/view"
)

func NewBoatView() IViewManager[domain.Boat, gin.H] {
	return view.BoatView{}
}
