🧩 Desafio de Golang (Versão Revisada e Organizada)
Foco: Structs, Interfaces, Composição e Responsabilidade Clara
🎯 Objetivo

Implementar um sistema simples de pagamentos, usando Go de forma idiomática, com:

Structs

Interfaces

Métodos

Composição (sem herança)

Polimorfismo real

Tratamento de erro

O foco é design limpo, não quantidade de código.

🧠 Princípios de Design (importante)

Wallet guarda estado (saldo e histórico)

Métodos de pagamento executam lógica, não armazenam histórico

Transação é um registro de fato ocorrido, não um serviço

User coordena a operação (caso de uso)

📌 Modelo Conceitual
User
 └── Wallet
      ├── Balance
      └── Transactions[]

User → usa → PaymentMethod (interface)
PaymentMethod ← CreditCard | Pix

📌 Requisitos Funcionais (Reescritos)
1️⃣ Usuário (User)

Crie uma struct User que represente um usuário do sistema.

O usuário deve conter:

ID

Nome

Email

Uma Wallet embutida (composição obrigatória)

Responsabilidade do User

Orquestrar a operação de pagamento

Não executar lógica específica de cartão ou PIX

2️⃣ Carteira (Wallet)

Crie uma struct Wallet responsável por:

Armazenar o saldo atual

Armazenar o histórico de transações

Regras

A carteira não existe sem o usuário

Apenas a carteira:

Atualiza saldo

Registra transações

3️⃣ Transação (Transaction)

Crie uma struct Transaction que represente um evento financeiro ocorrido.

Cada transação deve conter:

ID

Valor

Nome do método de pagamento

Status (SUCCESS ou FAILED)

Observação importante

👉 Transaction não possui métodos de negócio, é apenas um registro de dados.

4️⃣ Interface de Método de Pagamento

Crie uma interface PaymentMethod.

Ela deve obrigar qualquer implementação a:

type PaymentMethod interface {
    Pay(amount float64) error
    Name() string
}

Regras

Nenhum código deve conhecer o tipo concreto (CreditCard, Pix, etc)

Não pode haver if ou switch para identificar o método

5️⃣ Métodos de Pagamento (Implementações)
a) Cartão de Crédito (CreditCard)

A struct deve conter:

ID do cartão

Limite total

Limite disponível

Regras

Não permitir pagamento acima do limite disponível

Reduzir o limite disponível após pagamento bem-sucedido

b) PIX (Pix)

A struct deve conter:

Chave PIX

Limite diário disponível

Regras

Não permitir pagamento acima do limite diário

Reduzir o limite diário após pagamento bem-sucedido

6️⃣ Caso de Uso: Realizar Pagamento

Crie um método do User:

func (u *User) MakePayment(method PaymentMethod, amount float64) error

Fluxo correto:

Solicita o pagamento via interface (method.Pay)

Se falhar:

Retorna erro

NÃO altera saldo

NÃO registra transação

Se tiver sucesso:

Atualiza o saldo da Wallet

Cria uma Transaction com status SUCCESS

Registra a transação na Wallet

🧪 Cenários Obrigatórios no main

No main.go, demonstre:

Criação de usuário com saldo zero

Pagamento válido com cartão

Pagamento inválido com cartão (limite excedido)

Pagamento válido com PIX

Impressão:

Saldo final da carteira

Histórico de transações

📐 Restrições Técnicas

❌ Não usar herança
❌ Não usar switch ou if para identificar métodos de pagamento
❌ Não usar bibliotecas externas

✔️ Usar composição
✔️ Usar interfaces corretamente
✔️ Erros idiomáticos do Go (error)