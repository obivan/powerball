package ticket

import "math/rand"

const bound = 70
const lastBound = 27

// Cost is the cost of ticket
const Cost = 2

// Ticket represents a ticket
type Ticket [6]uint

// New return a new random ticket
func New(r *rand.Rand) Ticket {
	t := Ticket{}
	for i := 0; i < len(t)-1; i++ {
		rn := makeRandomNumber(r, bound)
		if !t.contains(rn) {
			t[i] = rn
		} else {
			t[i] = makeRandomNumberExcept(r, bound, rn)
		}
	}
	t[5] = makeRandomNumber(r, lastBound)
	return t
}

func makeRandomNumber(r *rand.Rand, bound int) uint {
	return 1 + uint(r.Intn(bound-1))
}

func makeRandomNumberExcept(r *rand.Rand, bound int, n uint) uint {
	rn := makeRandomNumber(r, bound)
	if rn != n {
		return rn
	}
	return makeRandomNumberExcept(r, bound, n)
}

func (t Ticket) contains(n uint) bool {
	for _, v := range t {
		if v == n {
			return true
		}
	}
	return false
}

// MatchCount returns count of match numbers
func (t Ticket) MatchCount(another Ticket) uint {
	var count uint
	for i := 0; i < len(t)-1; i++ {
		n1 := t[i]
		for j := 0; j < len(t)-1; j++ {
			n2 := another[j]
			if n1 == n2 {
				count++
			}
		}
	}
	return count
}

// MatchLast checks match of last number
func (t Ticket) MatchLast(another Ticket) bool {
	return t[len(t)-1] == another[len(t)-1]
}

// Prize returns the winnings
func (t Ticket) Prize(another Ticket) uint {
	var prize uint
	matchLast := t.MatchLast(another)
	matchCount := t.MatchCount(another)
	if matchLast {
		switch matchCount {
		case 0, 1:
			prize = 4
		case 2:
			prize = 7
		case 3:
			prize = 100
		case 4:
			prize = 50000
		case 5:
			prize = 1500000000
		}
	} else {
		switch matchCount {
		case 3:
			prize = 7
		case 4:
			prize = 100
		case 5:
			prize = 1000000
		}
	}
	return prize
}
