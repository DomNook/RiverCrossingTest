package event

import (
	"testing"

	"github.com/DomNook/RiverCrossingTest/event"
	"github.com/DomNook/RiverCrossingTest/state"
)

func TestPutFoxInBoat(t *testing.T) {
	event.PutFoxInBoat()
	wanted := true
	foxstate := state.FoxInBoat
	if foxstate != wanted {
		t.Errorf("Error, got %v, wanted %v.", foxstate, wanted)
	}
}
