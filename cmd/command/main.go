package main

import (
	"fmt"

	"github.com/0ndreu/patterns-wb/pkg/command/multiply"
	"github.com/0ndreu/patterns-wb/pkg/command/reciever"
	"github.com/0ndreu/patterns-wb/pkg/command/sum"
)

func main() {
	calculator := reciever.NewCalculator(2, 3)
	multiplyCommand := &multiply.MultiplyCommand{calculator}
	sumCom := &sum.SumCommand{calculator}
	sumCom.Execute()
	multiplyCommand.Execute()
	fmt.Println()
}
