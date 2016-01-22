package ticket

import (
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const randomCount = 100000

var r = rand.New(rand.NewSource(rand.Int63()))

func TestRandomNumber(t *testing.T) {
	for i := 0; i < randomCount; i++ {
		n := makeRandomNumber(r, bound)
		if n < 1 || n > bound {
			t.Fatalf("random number = %d", n)
		}
	}
}

func TestNew(t *testing.T) {
	for i := 0; i < randomCount; i++ {
		ti := New(r)

		if len(ti) < 6 {
			t.Fatal("Len of ticket < 6")
		}

		for i := 0; i < len(ti)-1; i++ {
			if ti[i] < 0 || ti[i] > bound {
				t.Fatalf("ticket t[n] should be 0 < n < 70 where n=1..5. t=%v", ti)
			}
		}

		if ti[5] < 0 || ti[5] > lastBound {
			t.Fatalf("ticket t[5] should be 0 < n < 27. t=%v", ti)
		}

	}
}

func TestMatchLast(t *testing.T) {
	t1 := Ticket{1, 2, 3, 4, 5, 6}
	t2 := Ticket{1, 2, 3, 4, 5, 6}
	if !t1.MatchLast(t2) {
		t.Fail()
	}
}

func TestMatchCount(t *testing.T) {
	t1 := Ticket{1, 2, 3, 4, 5, 6}
	t2 := Ticket{1, 2, 3, 4, 5, 6}
	if t1.MatchCount(t2) != 5 {
		t.Fail()
	}
	t1 = Ticket{1, 2, 3, 4, 5, 6}
	t2 = Ticket{5, 6, 7, 8, 9, 10}
	if t1.MatchCount(t2) != 1 {
		t.Fail()
	}
	t1 = Ticket{1, 2, 3, 4, 5, 6}
	t2 = Ticket{7, 8, 9, 10, 11, 12}
	if t1.MatchCount(t2) != 0 {
		t.Fail()
	}
}

func TestPrize(t *testing.T) {
	t1 := Ticket{1, 2, 3, 4, 5, 6}
	if t1.Prize(Ticket{7, 8, 9, 10, 11, 12}) != 0 {
		t.Fatal("prize must be 0")
	}
	if t1.Prize(Ticket{7, 8, 9, 10, 11, 6}) != 4 {
		t.Fatalf("prize must be 4")
	}
	if t1.Prize(Ticket{7, 8, 9, 10, 1, 6}) != 4 {
		t.Fatalf("prize must be 4")
	}
	if t1.Prize(Ticket{7, 8, 9, 2, 1, 6}) != 7 {
		t.Fatalf("prize must be 7")
	}
	if t1.Prize(Ticket{7, 8, 3, 2, 1, 6}) != 100 {
		t.Fatalf("prize must be 100")
	}
	if t1.Prize(Ticket{7, 4, 3, 2, 1, 6}) != 50000 {
		t.Fatalf("prize must be 50000")
	}
	if t1.Prize(Ticket{5, 4, 3, 2, 1, 6}) != 1500000000 {
		t.Fatalf("prize must be 1500000000")
	}
	if t1.Prize(Ticket{7, 8, 3, 2, 1, 20}) != 7 {
		t.Fatalf("prize must be 7")
	}
	if t1.Prize(Ticket{7, 4, 3, 2, 1, 20}) != 100 {
		t.Fatalf("prize must be 100")
	}
	if t1.Prize(Ticket{5, 4, 3, 2, 1, 20}) != 1000000 {
		t.Fatalf("prize must be 1000000")
	}
}
