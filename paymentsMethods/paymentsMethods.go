package paymentsMethods

type Pagamento interface {
	Processar() string
	Valor() float64
}

type Transacao struct {
	ValorPago float64
}
