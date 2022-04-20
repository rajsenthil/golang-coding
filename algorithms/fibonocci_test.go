package datastructures

import (
	"fmt"
	"testing"
)

func TestFibonocci(t *testing.T) {
	fib0 := fibonocci(0)
	fmt.Printf("Fibonocci 0: %v\n", fib0)

	fib1 := fibonocci(1)
	fmt.Printf("Fibonocci 1: %v\n", fib1)

	fib2 := fibonocci(2)
	fmt.Printf("Fibonocci 2: %v\n", fib2)

	fib3 := fibonocci(3)
	fmt.Printf("Fibonocci 3: %v\n", fib3)

	fib4 := fibonocci(4)
	fmt.Printf("Fibonocci 4: %v\n", fib4)

	fib50 := fibonocci(50)
	fmt.Printf("Fibonocci 50: %v\n", fib50)

	fib51 := fibonocci(51)
	fmt.Printf("Fibonocci 51: %v\n", fib51)

	fib60 := fibonocci(60)
	fmt.Printf("Fibonocci 60: %v\n", fib60)

	fib75 := fibonocci(75)
	fmt.Printf("Fibonocci 75: %v\n", fib75)

	fib80 := fibonocci(80)
	fmt.Printf("Fibonocci 80: %v\n", fib80)

	fib90 := fibonocci(90)
	fmt.Printf("Fibonocci 90: %v\n", fib90)

	fib92 := fibonocci(92)
	fmt.Printf("Fibonocci 92: %v\n", fib92)

	fib100 := fibonocci(100)
	fmt.Printf("Invalid: Due to overflow:  Fibonocci 100: %v\n", fib100)

}
