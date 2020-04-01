package credithistory

import (
	"math/rand"
	"time"

	"github.com/0ndreu/patterns-wb/pkg/facade/models/facade"
)

// CreditHistory checks last credit status and history
type CreditHistory interface {
	CheckCreditHistory(time.Time) bool
	CountOfCredits(time time.Time)
}

type creditHistory struct {
	lastCredit     bool
	countOfCredits int
	isValid        bool
}

// CheckCreditHistory approval
func (ch *creditHistory) CheckCreditHistory(time time.Time) (isValid bool) {
	rand.Seed(time.UnixNano())
	chanceOfCredit := rand.Intn(30)
	if ch.lastCredit && chanceOfCredit+ch.countOfCredits >= 10 {
		isValid = true
	} else {
		isValid = false
	}
	return
}

// CountOfCredits sets count of credits to the NewCreditHistory instance
func (ch *creditHistory) CountOfCredits(time time.Time) {
	rand.Seed(time.UnixNano())
	ch.countOfCredits = rand.Intn(20)
}

// NewCreditHistory instance
func NewCreditHistory() CreditHistory {
	return &creditHistory{
		lastCredit:     models.LastCredit,
		countOfCredits: 0,
		isValid:        false,
	}
}
