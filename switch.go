package main

import "fmt"
import "time"

func main(){
	i := 2
	fmt.Print("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2: 
		fmt.Println("two")
	case 3:	
		fmt.Println("one")
	}
	
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it is the weekend")
	default:
		fmt.Println("it is a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("it is before noon")
	default:
		fmt.Println("it is after noon")
	}
}
