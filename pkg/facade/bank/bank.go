package bank

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Bank calculates the percentage of credit
type Bank interface {
	GetPercentForCredit() float32
	BankLastCredit(time time.Time)
}

type bank struct {
	desireCredit uint
	lastCredit   uint
}

// GetPercentForCredit calculate size of percent per month for credit
func (b *bank) GetPercentForCredit() (percent float32) {
	var tmp float32

	if b.desireCredit <= b.lastCredit {
		tmp := float32(b.lastCredit / b.desireCredit)
		percent = tmp * 0.007
		if percent <= 0.055 {
			percent = 0.055
		} else if percent >= 0.08 {
			percent = 0.065
		}
	} else {
		tmp = float32(b.desireCredit / b.lastCredit)
		percent = tmp * 0.05
		if percent <= 0.07 {
			percent = 0.07
		} else if percent >= 0.1 {
			percent = 0.1
		}
	}
	return
}

// BankLastCredit sets sum of last credit
func (b *bank) BankLastCredit(time time.Time) {
	rand.Seed(time.Unix())
	minLastCr := rand.Intn(10000)
	maxLastCr := rand.Intn(1000000000)
	b.lastCredit = uint(math.Abs(float64(maxLastCr - minLastCr)))
}

// NewBank instance
func NewBank(desireCredit uint) (Bank, error) {
	if desireCredit == 0 {
		err := fmt.Errorf("you can't take zero credit")
		return nil, err
	}
	return &bank{
		desireCredit: desireCredit,
		lastCredit:   0,
	}, nil
}
