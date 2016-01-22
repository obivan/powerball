package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/obivan/powerball/ticket"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const ticketGeneratorsCount = 2

var randomTicketCh = make(chan ticket.Ticket, 1000)
var money uint

func startTicketGenerator() {
	for i := 0; i < ticketGeneratorsCount; i++ {
		go func() {
			r := rand.New(rand.NewSource(rand.Int63()))
			for {
				randomTicketCh <- ticket.New(r)
			}
		}()
	}
}

func buyTicket() (ticket.Ticket, error) {
	if money >= ticket.Cost {
		money = money - ticket.Cost
		return <-randomTicketCh, nil
	}
	return ticket.Ticket{}, fmt.Errorf("You have out of money")
}

func main() {
	// f, err := os.Create("cpu.profile")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()
	// pprof.StartCPUProfile(f)
	// defer pprof.StopCPUProfile()

	// f, err := os.Create("trace.trc")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()
	// trace.Start(f)
	// defer trace.Stop()

	flag.UintVar(&money, "m", 1000000, "your money")
	flag.Parse()
	fmt.Printf("Your money: %d\n", money)

	startTicketGenerator()

	ticket := <-randomTicketCh
	fmt.Printf("Your ticket: %2v\n", ticket)

	iters := 0
	for {
		next, err := buyTicket()
		if err != nil {
			fmt.Println(err)
			break
		}

		money += ticket.Prize(next)

		if ticket.Prize(next) == 1500000000 {
			fmt.Println(iters)
			fmt.Println("Jackpot! $1500000000")
			fmt.Printf("You won after %d iterations!\n", iters)
			fmt.Printf("Winner ticket is %v\n", next)
			fmt.Printf("match count: %d | match last: %v | prize: $%d\n",
				ticket.MatchCount(next), ticket.MatchLast(next), ticket.Prize(next))
			break
		}
		iters++
		if iters%1000000 == 0 {
			fmt.Print(".")
		}
		if iters%25000000 == 0 {
			fmt.Printf(" %3dM iters | your balance: $%d\n", iters/1000000, money)
		}
	}

	fmt.Println()
	fmt.Printf("Iterations count: %d\n", iters)
	fmt.Printf("Your balance: $%d\n", money)
}
