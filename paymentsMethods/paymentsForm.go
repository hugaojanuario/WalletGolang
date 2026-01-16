package paymentsMethods

import "fmt"

type CartaoCredito struct {
	Transacao
	Parcelas int
}

func (c *CartaoCredito) Processar() string {

	return fmt.Sprintln("Pagamento no cartão em", c.Parcelas, "parcelas")
}

func (c CartaoCredito) Valor() float64 {
	juros := 0.02 * float64(c.Parcelas)
	return c.ValorPago * (1 + juros)
}

type Pix struct {
	Transacao
}

func (p *Pix) Processar() string {
	return fmt.Sprintln("Pagamento via PIX")
}

func (p *Pix) Valor() float64 {
	return p.ValorPago
}

type Boleto struct {
	Transacao
	DiasAtraso int
}

func (b *Boleto) Processar() string {
	return fmt.Sprintln("Pagamento via Boleto")
}

func (b *Boleto) Valor() float64 {
	if b.DiasAtraso > 0 {
		juros := 0.01 * float64(b.DiasAtraso)
		return b.ValorPago * (1 + juros)
	} else {
		return b.ValorPago
	}

}

func ExecutarPagamento(p Pagamento) {
	fmt.Println(p.Processar())
	fmt.Println(p.Valor())

}
