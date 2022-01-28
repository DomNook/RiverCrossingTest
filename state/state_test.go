package state

import (
	"testing"
)

func TestFetchState(t *testing.T) {
	wanted := "currently onboard the boat:\n \ncurrently on starting land:\nChicken\nFox\nCorn\nFarmer\n\ncurretnly on the ending land:\n "
	state := state.FetchState()
	if state != wanted {
		t.Errorf("Error, got %q, wanted %q.", state, wanted)
	}
}
