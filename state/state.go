package state

//Variable for boat location
//where false = starting land
//and true = winning land
var WhereIsBoat = false

// Variables for verifiying if objects are in the boat
var ChickenInBoat = false
var FoxInBoat = false
var CornInBoat = false
var FarmerInBoat = false

// Variables for verifiying if objects are on the starting land
var ChickenOnOldLand = true
var FoxOnOldLand = true
var CornOnOldLand = true
var FarmerOnOldLand = true

// Variables for verifiyng if objects are on the winning land
var ChickenOnNewLand = false
var FoxOnNewLand = false
var CornOnNewLand = false
var FarmerOnNewLand = false

// Total slots in boat
var Boatslots = 4

var textboat = "\nThe Boat is currently at the: "
var boatlocation = ""

var text1 = "\ncurrently onboard the boat:\n"
var text2 = "\ncurrently on starting land:\n"
var text3 = "\ncurretnly on the ending land:\n"

var textsub1 = ""
var textsub2 = ""
var textsub3 = ""

func FetchState() string {
	//where is the boat?
	if WhereIsBoat == false {
		boatlocation = "starting land\n"
	} else {
		boatlocation = "winning land\n"
	}

	//are the objects in the boat?

	//first empty the string
	textsub1 = ""
	if ChickenInBoat == true {
		textsub1 = "Chicken\n"
	}
	if FoxInBoat == true {
		textsub1 = textsub1 + "Fox\n"
	}
	if CornInBoat == true {
		textsub1 = textsub1 + "Corn\n"
	}
	if FarmerInBoat == true {
		textsub1 = textsub1 + "Farmer\n"
	}

	//are the objects on the starting landmass?

	//first empty the string
	textsub2 = ""
	if ChickenOnOldLand == true {
		textsub2 = "Chicken\n"
	}
	if FoxOnOldLand == true {
		textsub2 = textsub2 + "Fox\n"
	}
	if CornOnOldLand == true {
		textsub2 = textsub2 + "Corn\n"
	}
	if FarmerOnOldLand == true {
		textsub2 = textsub2 + "Farmer\n"
	}

	//are the objects on the ending landmass?

	//first empty the string
	textsub3 = ""
	if ChickenOnNewLand == true {
		textsub3 = "Chicken\n"
	}
	if FoxOnNewLand == true {
		textsub3 = textsub3 + "Fox\n"
	}
	if CornOnNewLand == true {
		textsub3 = textsub3 + "Corn\n"
	}
	if FarmerOnNewLand == true {
		textsub3 = textsub3 + "Farmer\n"
	}

	// return previously fetched state values
	fulltext := textboat + boatlocation + text1 + textsub1 + text2 + textsub2 + text3 + textsub3
	return fulltext
}
