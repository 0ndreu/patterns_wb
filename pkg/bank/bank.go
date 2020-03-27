package bank

import (
	"math"
	"math/rand"
	"time"
)

type Bank interface {
	GetPercentForCredit() float32
	//GiveCredit(creditHistory bool) (float32, bool)
}

type bank struct {
	desireCredit uint
	lastCredit   uint
}

// size of percent per month for credit
func (b *bank) GetPercentForCredit() float32 {
	var (
		percent float32
		tmp     float32
	)
	if b.desireCredit <= b.lastCredit {
		tmp = float32(b.lastCredit / b.desireCredit)
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
	return percent
}

func NewBank(desireCredit uint) Bank {
	rand.Seed(time.Now().UnixNano())
	a := rand.Intn(10000)
	b := rand.Intn(1000000000)
	lastCredit := uint(math.Abs(float64(b - a)))
	return &bank{
		desireCredit: desireCredit,
		lastCredit:   lastCredit,
	}
}
