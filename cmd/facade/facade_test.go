package main

import (
	"fmt"
	"testing"
	"time"

	"patterns-wb/pkg/facade/bank"
	"patterns-wb/pkg/facade/credithistory"
	"patterns-wb/pkg/facade/facade"
)

var okResult = `Dear Petrov
you can't take credit in my bank.`

func TestOk(t *testing.T) {
	var sum uint
	name := "Petrov"
	sum = 500000
	time := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	ch := creditHistory.NewCreditHistory()
	ch.CountOfCredits(time)
	b, err := bank.NewBank(sum)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	b.BankLastCredit(time)
	credit := facade.NewCredit(b, ch)
	out := credit.GiveCredit(name, ch, b, time)
	if out != okResult {
		t.Errorf(out)
	}
}

func TestZeroSum(t *testing.T) {
	var sum uint
	sum = 0
	time := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	ch := creditHistory.NewCreditHistory()
	ch.CountOfCredits(time)
	_, err := bank.NewBank(sum)
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Errorf("Failed zero sum test")
}
