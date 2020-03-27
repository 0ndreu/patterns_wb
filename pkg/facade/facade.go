package facade

import (
	"fmt"
)

type Bank interface {
	GetPercentForCredit() float32
}

type CreditHistory interface {
	CheckCreditHistory() bool
}

type Credit interface {
	GiveCredit(string, CreditHistory, Bank)
}

type credit struct {
	bank       Bank
	creditHist CreditHistory
}

// percent for credit and credit approval status
func (c *credit) getCredit(ch CreditHistory, b Bank) (float32, bool) {
	approval := ch.CheckCreditHistory()
	var percentForCredit float32
	if approval {
		percentForCredit = b.GetPercentForCredit()
	}
	return percentForCredit, approval
}

func (c *credit) GiveCredit(user string, ch CreditHistory, b Bank) {
	percent, isApprove := c.getCredit(ch, b)
	c.returnInfo(user, isApprove, percent)
}

func (c *credit) returnInfo(name string, isApprove bool, percent float32) {
	fmt.Println("Dear", name)
	if isApprove {
		fmt.Println("you can take credit in my bank with", percent*100, "% per month.")
	} else {
		fmt.Println("you can't take credit in my bank.")
	}
}

func NewCredit() Credit {
	return &credit{
		bank:       nil,
		creditHist: nil,
	}
}
