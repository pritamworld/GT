package funding
import(
	"fmt"
)
type FundServer struct {
    commands chan interface{}
    fund Fund
}

type WithdrawCommand struct {
    Amount int
}

type BalanceCommand struct {
    Response chan int
}

func (s *FundServer) Balance() int {
    responseChan := make(chan int)
    s.commands <- BalanceCommand{ Response: responseChan }
    return <- responseChan
}

func (s *FundServer) Withdraw(amount int) {
    s.commands <- WithdrawCommand{ Amount: amount }
}

func NewFundServer(initialBalance int) *FundServer {
    server := &FundServer{
        // make() creates builtins like channels, maps, and slices
        commands: make(chan interface{}),
        fund: *NewFund(initialBalance),
        //Add * in front of *NewFund(initialBalance)
    }

    // Spawn off the server's main loop immediately
    go server.loop()
    return server
}

func (s *FundServer) loop() {
    for command := range s.commands {

        // command is just an interface{}, but we can check its real type
        switch command.(type) {

        case WithdrawCommand:
            // And then use a "type assertion" to convert it
            withdrawal := command.(WithdrawCommand)
            s.fund.Withdraw(withdrawal.Amount)

        case BalanceCommand:
            getBalance := command.(BalanceCommand)
            balance := s.fund.Balance()
            getBalance.Response <- balance

        default:
            panic(fmt.Sprintf("Unrecognized command: %v", command))
        }
    }
}

//https://www.toptal.com/go/go-programming-a-step-by-step-introductory-tutorial