package main

import (
	"fmt"
	"time"

	"patterns-wb/pkg/facade/bank"
	"patterns-wb/pkg/facade/credithistory"
	"patterns-wb/pkg/facade/facade"
)

func main() {
	var sum uint
	name := "Petrov"
	sum = 500000
	time := time.Now()
	ch := credithistory.NewCreditHistory()
	ch.CountOfCredits(time)
	b, err := bank.NewBank(sum)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	b.BankLastCredit(time)
	credit := facade.NewCredit(b, ch)
	out := credit.GiveCredit(name, ch, b, time)
	fmt.Println(out)
}
