package state

import (
	"testing"

	"github.com/DomNook/RiverCrossingTest/state"
)

func TestFetchState(t *testing.T) {
	wanted := "\nThe Boat is currently at the: starting land\n\ncurrently onboard the boat:\n\ncurrently on starting land:\nChicken\nFox\nCorn\nFarmer\n\ncurretnly on the ending land:\n"
	state := state.FetchState()
	if state != wanted {
		t.Errorf("Error, got %q, wanted %q.", state, wanted)
	}
}
