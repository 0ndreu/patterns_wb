package facade

import (
	"fmt"
	"time"
)

type bank interface {
	GetPercentForCredit() float32
}

type creditHistory interface {
	CheckCreditHistory(time.Time) bool
}

// Interface to start work with facade
type Credit interface {
	GiveCredit(string, creditHistory, bank, time.Time) (string, error)
}

type credit struct {
	bank       bank
	creditHist creditHistory
}


// GiveCredit is entry point to facade
func (c *credit) GiveCredit(user string, ch creditHistory, b bank, g time.Time) (out string, err error) {
	if user != ""{
		percent, isApprove := c.getCredit(ch, b, g)
		out = c.returnInfo(user, isApprove, percent)
		return
	}
	err = fmt.Errorf("null name")
	return "", nil
}

// getCredit calculate percent for credit and credit approval status
func (c *credit) getCredit(ch creditHistory, b bank, g time.Time) (float32, bool) {
	approval := ch.CheckCreditHistory(g)
	var percentForCredit float32
	if approval {
		percentForCredit = b.GetPercentForCredit()
	}
	return percentForCredit, approval
}

// returnInfo returns string retrieve
func (c *credit) returnInfo(name string, isApprove bool, percent float32) (out string) {
	out = "Dear " + name
	totalPercent := fmt.Sprintf("%.2f", percent*100)
	if isApprove {
		out += "\nyou can take credit in my bank with " + totalPercent + "% per month."
	} else {
		out += "\nyou can't take credit in my bank."
	}
	return
}

// NewCredit instance
func NewCredit(bank bank, creditHist creditHistory) Credit {
	return &credit{
		bank:       bank,
		creditHist: creditHist,
	}
}
