package main

import (
	"patterns_wb/pkg/bank"
	"patterns_wb/pkg/creditHistory"
	"patterns_wb/pkg/facade"
)

func main() {
	name := "Petrov"
	var sum uint
	sum = 100000000
	credit := facade.NewCredit()
	ch := creditHistory.NewCreditHistory()
	b := bank.NewBank(sum)

	credit.GiveCredit(name, ch, b)
}
