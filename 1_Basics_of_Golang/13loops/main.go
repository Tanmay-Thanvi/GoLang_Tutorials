package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("----Loops----")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday"}

	fmt.Println(days)

	// there is only one loop : for but have many variation
	for d := 0; d < len(days); d++ { // there is nothing like ++d
		fmt.Println(days[d])
	}

	for i := range days {
		fmt.Println(i, days[i]) // i gives index here
	}

	for index, day := range days {
		fmt.Println(index, day)
	}

	roguevalue := 1

	for roguevalue < 10 {

		if roguevalue == 7 {
			break
		}

		if roguevalue == 5 {
			roguevalue++
			continue
		}

		if roguevalue == 2 {
			goto lco
		}

		fmt.Println("Value : ", roguevalue)
		roguevalue++
	}

lco:
	fmt.Println("Jumping at lco.dev")

	// infinite loop in golang (Asked in Pubmatic)
	cnt := 0
	for {
		cnt += 5
		if cnt > 50 {
			break;
		}
		fmt.Println("i = ", cnt)
		time.Sleep(5)
	}
}
