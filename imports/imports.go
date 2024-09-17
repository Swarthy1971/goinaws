package imports

import (
	"fmt"
)

type Ticket struct {
	ID int
	// event string	// lowercase event can't be used outside the package
	Event string
}

func (t Ticket) PrintEvent() { // method with lower	case can't be used outside the package
	fmt.Println(t.Event)
}
