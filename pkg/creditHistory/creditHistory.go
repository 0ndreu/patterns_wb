package creditHistory

import (
	"math/rand"
	"time"

	"patterns_wb/pkg/models/facade"
)

// CreditHistory checks last credit status and history
type CreditHistory interface {
	CheckCreditHistory() bool
}

type creditHistory struct {
	lastCredit     bool
	countOfCredits int
}

// Check history and approval
func (ch *creditHistory) CheckCreditHistory() bool {
	rand.Seed(time.Now().UnixNano())
	chanceOfCredit := rand.Intn(30)
	if ch.lastCredit && chanceOfCredit+ch.countOfCredits >= 10 {
		return true
	}
	return false
}

// NewCreditHistory instance
func NewCreditHistory() CreditHistory {
	rand.Seed(time.Now().UnixNano())
	countOfCredits := rand.Intn(20)
	return &creditHistory{
		lastCredit:     models.LastCredit,
		countOfCredits: countOfCredits,
	}
}
