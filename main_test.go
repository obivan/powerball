package main

import "testing"

func init() {
	startTicketGenerator()
}

func BenchmarkGenCh(b *testing.B) {
	for i := 0; i < b.N; i++ {
		<-randomTicketCh
	}
}
