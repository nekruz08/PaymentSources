package transfer

import "bank/pkg/bank/types"

const bonusPercent = 0.0050

// Bonus рассчитывает по amount начисляемый при переводе бонус.
func Bonus(amount types.Money) types.Money {
	bonus:=types.Money(float64(amount)*bonusPercent)
	return bonus
}

// Total рассчитывает по amount итоговую сумму для зачисления с учётом бонуса.
func Total(amount types.Money) types.Money {
	total:=amount + Bonus(amount)
	return	total
}