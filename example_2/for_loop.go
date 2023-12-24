package for_loop

import "fmt"

func Runner() {

	i := 1
	for i <= 3 {
		fmt.Print(i," ")
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Print(j, " ")
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Print(n," ")
	}
	if num := -9; num < 0 {
        fmt.Println("number ",num, "is negative")
    } else if num < 10 {
        fmt.Println("number ",num, "has 1 digit")
    } else {
        fmt.Println("number ",num, "has multiple digits")
    }

}