package payment

import (
	"bank/pkg/bank/types"
	"fmt"
)
func ExampleMax() {
	payments:=[]types.Payment{
		{
			ID: 1,
			Amount: 99,
		},
		{
			ID: 2,
			Amount: 100,
		},
		{
			ID: 3,
			Amount: 243,
		},
	}
	
	max:=Max(payments)
	fmt.Println(max.ID)
	//Output:3
	
}