package booksy

import (
	"strconv"
	"time"
)

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

type FIAT_DENOMINATION int32

func (fd FIAT_DENOMINATION) String() string {
	switch fd {
	case UNIT:
		return "UNIT"
	case SUB_UNIT:
		return "SUB_UNIT"
	default:
		return "UNKNOW"
	}
}

const (
	UNIT     FIAT_DENOMINATION = iota + 1 // representation of the currency in it's main unit e.x `$1`
	SUB_UNIT                              // representation of the currency in it's sub unit e.x `100 cents`
)

type TRANSACTION_STATE int32

func (ts TRANSACTION_STATE) String() string {
	switch ts {
	case SCHEDULED:
		return "SCHEDULED"
	case PENDING:
		return "PENDING"
	case COMPLETED:
		return "COMPLETED"
	case DECLINED:
		return "DECLINED"
	default:
		return "UNKNOW"
	}
}

const (
	SCHEDULED TRANSACTION_STATE = iota + 1
	PENDING
	COMPLETED
	DECLINED
)

type Transaction struct {
	ItemName        string
	Description     string
	Amount          int32 // transactional amount in cents
	Created         time.Time
	Updated         time.Time
	Deleted         time.Time
	Type            TRANSACTION_TYPE
	TypeDescription string
	ById            string // who initiated the transaction unique id
	By              string // who initidated the transaction full name
	State           TRANSACTION_STATE
}

// Returns [Transaction] amount in it's [UNIT] format
//
// For exmaple 100 cents = 1.00
func (t Transaction) AmountInDollars() string {
	result := float64(t.Amount / 100)
	return strconv.FormatFloat(float64(result), 'f', 2, 64)
}
