package creditHistory

import (
	"math/rand"
	"patterns_wb/pkg/models/facade"
	"time"
)

type CreditHistory interface {
	CheckCreditHistory() bool
}

type creditHistory struct {
	lastCredit     bool
	countOfCredits int
}

// check history and approval
func (ch *creditHistory) CheckCreditHistory() bool {
	rand.Seed(time.Now().UnixNano())
	chanceOfCredit := rand.Intn(30)
	if ch.lastCredit && chanceOfCredit+ch.countOfCredits >= 10 {
		return true
	}
	return false
}

func NewCreditHistory() CreditHistory {
	rand.Seed(time.Now().UnixNano())
	countOfCredits := rand.Intn(20)
	return &creditHistory{
		lastCredit:     models.LastCredit,
		countOfCredits: countOfCredits,
	}
}
