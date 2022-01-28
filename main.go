package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/DomNook/RiverCrossingTest/event"
	"github.com/DomNook/RiverCrossingTest/state"
)

func main() {
	fmt.Println("The beginning")
	fmt.Println(state.FetchState())

	fmt.Println("What would you like to do? [Insert to boat [Press 1]] [View State [Press 2]] \n [Send the boat across [Press 3]]")
	reader := bufio.NewReader(os.Stdin)

	char, _, err := reader.ReadRune()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(char)

	if char == 49 {
		fmt.Println("Under construction")
	} else if char == 50 {
		fmt.Println(state.FetchState())
	} else if char == 51 {
		fmt.Println("if i had implemented the boat sailing function the boat would have went")
	} else {
		fmt.Println("Error: Inputted number is not allowed. Program will now exit")
	}

}

func ChooseMoveItem() {

	reader := bufio.NewReader(os.Stdin)

	char, _, err := reader.ReadRune()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(char)

	if char == 49 {
		event.PutChickenInBoat()
		fmt.Println(state.FetchState())
	} else if char == 50 {
		event.PutFoxInBoat()
		fmt.Println(state.FetchState)
	} else if char == 51 {
		event.PutCornInBoat()
		fmt.Println(state.FetchState)
	} else if char == 52 {
		event.PutFarmerInBoat()
		fmt.Println(state.FetchState)
	} else {
		fmt.Println("Error: Inputted number is not allowed. Program will now exit")
	}
}
