package tourPackage

import (
	"fmt"
	"time"
)

func SwitchCase() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("today")
	case today + 1:
		fmt.Println("tomorrow")
	case today + 0:
		fmt.Println("in 2 days")
	default:
		fmt.Println("too far away")
	}
}
