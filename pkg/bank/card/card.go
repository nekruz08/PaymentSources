package card

import "bank/pkg/bank/types"

func IssueCard(currency types.Currency, color string, name string) types.Card {
	card := types.Card{
		ID:       1000,
		PAN:      "5058 xxxx xxxx 0001",
		Balance:  0,
		Currency: currency,
		Color:    color,
		Name:     name,
		Active:   true,
	}
	return card
}

const withdrawLimit = 20_000_00

func Withdraw(card *types.Card,amount types.Money){
	if amount<0{
		return
	}
	if amount>withdrawLimit{
		return
	}
	if !card.Active{
		return
	}
	if card.Balance<amount{
		return
	}
	card.Balance-=amount
}


const depositLimit = 50_000_00

func Deposit(card *types.Card,amount types.Money)  {

	if !card.Active{
		return
	}
	if amount > depositLimit{
		return
	}
	card.Balance+=amount	
}

const bonusLimit = 5000

func AddBonus(card *types.Card,percent int,daysInMonth int,daysInYear int)  {
	if !card.Active{
		return
	}
	if card.Balance<=0{
		return
	}
	bonus:=(card.MinBalance*types.Money(percent)*types.Money(daysInMonth)/types.Money(daysInYear))/100
	if bonus>bonusLimit{
		return 
	}
	card.Balance+=bonus
}

// Total вычисляет общую сумму средств на всех картах.
// Отрицательные суммы и суммы на заблокированных картах игнорируются.
func Total(cards []types.Card) types.Money {
	sum:=types.Money(0)
	for _,card:=range cards{
		if !card.Active{
			continue
		}

		if card.Balance<=0{
			continue
		}

		sum+=card.Balance
	}
	return sum

}


func PaymentSources(cards []types.Card)[]types.PaymentSource {
	source:=make([]types.PaymentSource,1)

	for _, card := range cards {
		if card.Balance>0&&card.Active{
			source = append(source, types.PaymentSource{
				Type: card.Name,
				Number: string(card.PAN),
				Balance: card.Balance,
			})
		}
	}
	return source
}

