package tourPackage

import (
	"fmt"
	"time"
)

func SwitchNoCondition() {
	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("Morning")
	case t.Hour() < 17:
		fmt.Println("Morning")
	default:
		fmt.Println("Good evening")
	}
}
