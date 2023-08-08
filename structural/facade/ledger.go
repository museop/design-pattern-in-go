package facade

import "fmt"

type Ledger struct{}

func (s *Ledger) makeEntry(accountID, txnType string, amount int) {
	fmt.Printf("Make ledger entry for accountID %s with txnType %s for ammount %d\n", accountID, txnType, amount)
	return
}
