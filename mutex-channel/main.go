package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main() {
	var bankBalance int
	var balance sync.Mutex

	fmt.Printf("initial account balance: $%d.00\n", bankBalance)

	incomes := []Income{
		{Source: "Gaji", Amount: 500},
		{Source: "Freelance", Amount: 50},
		{Source: "Kontrakan", Amount: 100},
		{Source: "Bonus", Amount: 10},
	}

	wg.Add(len(incomes))

	for i, income := range incomes {

		go func(i int, income Income) {

			defer wg.Done()

			for week := 1; week <= 52; week++ {

				balance.Lock()

				bankBalance += income.Amount

				balance.Unlock()

				fmt.Printf("on week %d, your earned $%d.00 from %s\n", week, income.Amount, income.Source)
			}
		}(i, income)
	}

	wg.Wait()

	fmt.Printf("final bank balance: $%d.00\n", bankBalance)
}
