package stratedy

import "testing"

func TestNewPayment(t *testing.T) {
	a := NewPayment(18, "Muby", &Cash{})
	b := NewPayment(9090, "Luke", &Bank{})

	a.Pay()
	b.Pay()
}
