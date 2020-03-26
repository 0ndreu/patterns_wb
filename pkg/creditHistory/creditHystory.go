package creditHistory

import (
	"math/rand"
	"time"
)

type CreditHistory interface {
	Check() bool
}

type creditHystory struct {
	name           string
	lastCredit     bool
	countOfCredits int
}

// loan approval check
func (ch *creditHystory) CheckCreditHystory() bool {
	rand.Seed(time.Now().UnixNano())
	chanceOfCredit := rand.Intn(100)
	if ch.lastCredit {
		if chanceOfCredit+ch.countOfCredits >= 20 {
			return true
		}
	} else {
		if chanceOfCredit+ch.countOfCredits < 20 {
			return false
		}
	}
	return false
}
