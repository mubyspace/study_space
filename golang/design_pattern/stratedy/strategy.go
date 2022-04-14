package stratedy

import "fmt"

/*
在实际应用中， 我们对不同的场景要采取不同的应对措施，也就是不同的策略。定义一个接口，传入的不同对象实现了接口，因此自动调用对应的策略

定义了一系列算法，并将每个算法封装起来，使它们可以相互替换，且算法的改变不会影响使用算法的客户。
*/

type Payment struct {
	payctx   *PayCtx
	strategy Strategy
}

type Strategy interface {
	Pay(*PayCtx)
}

type PayCtx struct {
	money int
	name  string
}

func NewPayment(money int, name string, strategy Strategy) *Payment {
	return &Payment{
		payctx: &PayCtx{
			money: money,
			name:  name,
		},
		strategy: strategy,
	}
}

func (p *Payment) Pay() {
	p.strategy.Pay(p.payctx)
}

type Cash struct{}

func (cash *Cash) Pay(ctx *PayCtx) {
	fmt.Printf("%s使用现金支付了%d\n", ctx.name, ctx.money)
}

type Bank struct{}

func (bank *Bank) Pay(ctx *PayCtx) {
	fmt.Printf("%s使用现金支付了%d\n", ctx.name, ctx.money)
}
