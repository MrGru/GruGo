package main

import (
	"fmt"

	"github.com/MrGru/GruGo/cookbook/data_conversion_and_composition/currency"
)

func main() {
	userInput := "15.93"
	pennies, err := currency.ConvertStringDollarsToPennies(userInput)
	if err != nil {
		panic(err)
	}
	fmt.Printf("User input converted to %d pennies\n", pennies)
	pennies += 15
	dollars := currency.ConvertPenniesToDollarsString(pennies)
	fmt.Printf("Added 15 cents, new value is %s dollars\n", dollars)
}
