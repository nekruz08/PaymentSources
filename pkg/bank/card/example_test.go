package card

import (
	"bank/pkg/bank/types"
	"fmt"
)
func ExampleWithdraw() {
	card:=types.Card{Balance: 20_000_00,Active: true}
	Withdraw(&card,10_000_00)

	fmt.Println(card.Balance)
	//Output: 1000000
	
}

func ExampleDeposit_positive() {
	card:=types.Card{Balance: 0,Active: true}
	Deposit(&card,10_000_00)
	fmt.Println(card.Balance)
	//Output: 1000000
}

func ExampleDeposit_inactive() {
	card:=types.Card{Balance: 0,Active: false}
	Deposit(&card,10_000_00)
	fmt.Println(card.Balance)
	//Output:0
}

func ExampleDeposit_limit() {
	card:=types.Card{Balance: 0,Active: true}
	Deposit(&card,5000000)
	fmt.Println(card.Balance)
	//Output: 5000000
	
}

func ExampleAddBonus() {
	card:=types.Card{Active: true,MinBalance: 10_000,Balance: 10_000}
	AddBonus(&card,3,30,365)
	fmt.Println(card.Balance)
	//Output:10024
}

func ExampleAddBonus_inactive() {
	card:=types.Card{Active: false,MinBalance: 10_000,Balance: 10_000}
	AddBonus(&card,3,30,365)
	fmt.Println(card.Balance)
	//Output:10000
}

func ExampleAddBonus_negative() {
	card:=types.Card{Active: true,MinBalance: 10_000,Balance: -10_000}
	AddBonus(&card,3,30,365)
	fmt.Println(card.Balance)
	//Output:-10000
}

func ExampleAddBonus_limit() {
	card:=types.Card{Active: true,MinBalance: 2100_000,Balance: 10_000}
	AddBonus(&card,3,30,365)
	fmt.Println(card.Balance)
	//Output:10000
}

func ExampleTotal() {
	fmt.Println(Total([]types.Card{
		{
			Balance:1_000_00,
			Active: true,
		},
	}))
	fmt.Println(Total([]types.Card{
		{
			Balance:1_000_00,
			Active: true,
		},
		{
			Balance:2_000_00,
			Active: true,
		},
	}))
	fmt.Println(Total([]types.Card{
		{
			Balance: 1_000_00,
			Active: false,
		},
	}))
	fmt.Println(Total([]types.Card{
		{
			Balance: -1_000_00,
			Active: true,
		},
	}))

	// Output:
	// 100000
	// 300000
	// 0
	// 0
	
}

func ExamplePaymentSources() {
	cards:=[]types.Card{
		{
			PAN: "5058 xxxx xxxx 0001",
			Balance:0,
			Active: true,
		},
		{
			PAN: "5058 xxxx xxxx 0002",
			Balance:2_000_00,
			Active: false,
		},
		{
			PAN: "5058 xxxx xxxx 0003",
			Balance:-100,
			Active: true,
		},
		{
			PAN: "5058 xxxx xxxx 0004",
			Balance:1,
			Active: true,
		},
		{
			PAN: "5058 xxxx xxxx 0005",
			Balance:2,
			Active: true,
		},
		{	PAN: "5058 xxxx xxxx 0006",
			Balance:1000000,
			Active: true,
		},
	}
	sources:=PaymentSources(cards)
	for _, source := range sources {
		fmt.Printf("%+v\n",source.Number)
	}
	//Output:
	//5058 xxxx xxxx 0004
	//5058 xxxx xxxx 0005
	//5058 xxxx xxxx 0006
}


