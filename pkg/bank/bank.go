package bank

type Balance interface {
	getPercentForCredit()
	GiveCredit(creditHistory bool) (float32, bool)
}

// percent for credit and credit approval status
func GiveCredit(creditHistory bool) (float32, bool) {
	var percentForCredit float32
	if creditHistory == true {
		percentForCredit = getPercentForCredit()
	}
	return percentForCredit, creditHistory
}

// size of percent per month for credit
func getPercentForCredit() float32 {
	var (
		percent float32
		tmp     float32
	)
	if desireCredit <= lastCredit {
		tmp = lastCredit / desireCredit
		percent = tmp * 0.07
		if percent <= 0.055 {
			percent = 0.055
		} else if percent >= 0.65 {
			percent = 0.65
		}
	} else {
		tpm = desireCredit / lastCredit
		percent = tmp * 0.05
		if percent <= 0.07 {
			percent = 0.07
		} else if percent >= 0.1 {
			percent = 0.1
		}
	}
	return percent
}
