package booksy

type TRANSACTION_TYPE int32

func (tt TRANSACTION_TYPE) String() string {
	switch tt {
	case DEBIT:
		return "DEBIT"
	case CREDIT:
		return "CREDIT"
	default:
		return "UNKNOW"
	}
}

const (
	DEBIT  TRANSACTION_TYPE = iota + 1 // money goes into account
	CREDIT                             // money goes out of the account
)

type Transaction struct {
}
