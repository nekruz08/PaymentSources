package deposit

import "fmt"

func ExampleCalculate() {
	fmt.Println(Calculate(0,"TJS"))
	fmt.Println(Calculate(0,"USD"))
	fmt.Println(Calculate(500_000_00,"TJS"))
	fmt.Println(Calculate(500_000_00,"USD"))
	fmt.Println(Calculate(1000_000_00,"TJS"))
	fmt.Println(Calculate(1000_000_00,"USD"))

	//Output:
	//0 0
	//0 0
	//166666 250000
	//83333 125000
	//333333 500000
	//166666 250000
	
}