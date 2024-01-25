package goroutines

var (
	deposits  = make(chan int)
	balances  = make(chan int)
	withdraws = make(chan withdraw)
)

type withdraw struct {
	ch     chan bool
	amount int
}

func Withdraw(amount int) bool {
	ch := make(chan bool)
	withdraws <- withdraw{ch, amount}
	return <-ch
}

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func teller() {
	balance := 0
	select {
	case amount := <-deposits:
		balance += amount
	case balances <- balance:
	case with := <-withdraws:
		if with.amount > balance {
			with.ch <- false
		} else {
			balance -= with.amount
			with.ch <- true
		}
	}
}

func StartBank() {
	go teller()
}
