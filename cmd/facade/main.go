package main

import (
	"fmt"

	"patterns_wb/pkg/bank"
	"patterns_wb/pkg/creditHistory"
	"patterns_wb/pkg/facade"
)

func main() {
	var sum uint
	name := "Petrov"
	sum = 999
	credit := facade.NewCredit()
	ch := creditHistory.NewCreditHistory()
	b := bank.NewBank(sum)
	fmt.Println(b)
	credit.GiveCredit(name, ch, b)
}
