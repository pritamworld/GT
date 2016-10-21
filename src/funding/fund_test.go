/*
package funding

import "testing"

func BenchmarkFund(b *testing.B) {
    // Add as many dollars as we have iterations this run
    fund := NewFund(b.N)

    // Burn through them one at a time until they are all gone
    for i := 0; i < b.N; i++ {
        fund.Withdraw(1)
    }

    if fund.Balance() != 0 {
        b.Error("Balance wasn't zero:", fund.Balance())
    }
}

//To run this test open terminal, first set GOPATH is not set and then go test -bench . funding
*/
package funding

import (
    "sync"
    "testing"
)

const WORKERS = 10

/*
func BenchmarkWithdrawals(b *testing.B) {
    // Skip N = 1
    if b.N < WORKERS {
        return
    }

    // Add as many dollars as we have iterations this run
    fund := NewFund(b.N)

    // Casually assume b.N divides cleanly
    dollarsPerFounder := b.N / WORKERS

    // WaitGroup structs don't need to be initialized
    // (their "zero value" is ready to use).
    // So, we just declare one and then use it.
    var wg sync.WaitGroup

    for i := 0; i < WORKERS; i++ {
        // Let the waitgroup know we're adding a goroutine
        wg.Add(1)
        
        // Spawn off a founder worker, as a closure
        go func() {
            // Mark this worker done when the function finishes
            defer wg.Done()

            for i := 0; i < dollarsPerFounder; i++ {
                fund.Withdraw(1)
            }
            
        }() // Remember to call the closure!
    }

    // Wait for all the workers to finish
    wg.Wait()

    if fund.Balance() != 0 {
        b.Error("Balance wasn't zero:", fund.Balance())
    }
}
*/
func BenchmarkWithdrawals(b *testing.B) {
    // Skip N = 1
    if b.N < WORKERS {
        return
    }
     // Add as many dollars as we have iterations this run
    //fund := NewFund(b.N)

    server := NewFundServer(b.N)

	// Casually assume b.N divides cleanly
    dollarsPerFounder := b.N / WORKERS
	var wg sync.WaitGroup
    

    // Spawn off the workers
    for i := 0; i < WORKERS; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for i := 0; i < dollarsPerFounder; i++ {
                //server.commands <- WithdrawCommand{ Amount: 1 }
                // Stop when we're down to pizza money
	            if server.Balance() <= 10 {
	                break
	            }
	            server.Withdraw(1)
            }
        }()
    }

    // Wait for all the workers to finish
    wg.Wait()

    balanceResponseChan := make(chan int)
    server.commands <- BalanceCommand{ Response: balanceResponseChan }
    //balance := <- balanceResponseChan
	balance := server.Balance()
    if balance != 10 {
        //b.Error("Balance wasn't zero:", balance)
        b.Error("Balance wasn't ten dollars:", balance)
    }
}

//To run GOMAXPROCS=4 go test -bench . funding
