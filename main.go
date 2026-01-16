package main

import (
	"payments/paymentsMethods"
)

func main() {

	Cc := paymentsMethods.CartaoCredito{paymentsMethods.Transacao{200}, 10}

	Pix := paymentsMethods.Pix{paymentsMethods.Transacao{100}}

	Boleto := paymentsMethods.Boleto{paymentsMethods.Transacao{300}, 5}

	var pagamentos []paymentsMethods.Pagamento
	pagamentos = append(pagamentos, &Cc, &Pix, &Boleto)

	for _, pagamento := range pagamentos {
		paymentsMethods.ExecutarPagamento(pagamento)
	}

}
