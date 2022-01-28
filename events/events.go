package events

import (
	"github.com/DomNook/RiverCrossingTest/state"
)

//functions to put objects into boat
func PutChickenInBoat() {
	state.ChickenInBoat = true
	state.ChickenOnNewLand = false
	state.ChickenOnOldLand = false
}
func PutFoxInBoat() {
	state.FoxInBoat = true
	state.FoxOnNewLand = false
	state.ChickenOnOldLand = false
}
func PutCornInBoat() {
	state.CornInBoat = true
	state.CornOnNewLand = false
	state.CornOnOldLand = false
}
func PutFarmerInBoat() {
	state.FarmerInBoat = true
	state.FarmerOnNewLand = false
	state.FarmerOnOldLand = false
}

//functions to offload objects to winning land
func DumpChickenToWLand() {
	state.ChickenInBoat = false
	state.ChickenOnNewLand = true
	state.ChickenOnOldLand = false
}
func DumpFoxToWLand() {
	state.FoxInBoat = false
	state.FoxOnNewLand = true
	state.FoxOnOldLand = false
}
func DumpCornToWLand() {
	state.CornInBoat = false
	state.CornOnNewLand = true
	state.CornOnOldLand = false
}
func DumpFarmerToWLand() {
	state.FarmerInBoat = false
	state.FarmerOnNewLand = true
	state.FarmerOnOldLand = false
}

//functions to offload objects to starting land
func DumpChickenToSLand() {
	state.ChickenInBoat = false
	state.ChickenOnNewLand = false
	state.ChickenOnOldLand = true
}
func DumpFoxToSLand() {
	state.FoxInBoat = false
	state.FoxOnNewLand = false
	state.FoxOnOldLand = true
}
func DumpCornToSLand() {
	state.CornInBoat = false
	state.CornOnNewLand = false
	state.CornOnOldLand = true
}
func DumpFarmerToSLand() {
	state.FarmerInBoat = false
	state.FarmerOnNewLand = false
	state.FarmerOnOldLand = true
}
