🧩 Desafio de Golang — Structs, Interfaces e Composição
🎯 Objetivo

Desenvolver um sistema simples de pagamentos, utilizando os conceitos fundamentais de programação orientada a objetos em Go, com foco em:

Structs

Interfaces

Métodos

Composição (structs aninhadas)

Tratamento de erros

📌 Requisitos Funcionais
1️⃣ Usuário

Crie uma estrutura que represente um usuário do sistema.

O usuário deve possuir:

Um identificador único

Nome

E-mail

Uma carteira financeira, que deve ser representada por outra estrutura

A carteira não deve existir separada do usuário (use composição).

2️⃣ Carteira (Wallet)

A carteira é responsável por:

Armazenar o saldo atual

Manter o histórico de transações realizadas

Cada transação deve ser registrada dentro da carteira do usuário.

3️⃣ Interface de Pagamento

Crie uma interface que represente um método de pagamento genérico.

Essa interface deve obrigar qualquer implementação a:

Executar um pagamento dado um valor

Informar o nome do método de pagamento utilizado

O sistema não pode conhecer o tipo concreto do método de pagamento (polimorfismo obrigatório).

4️⃣ Métodos de Pagamento

Implemente dois tipos diferentes de pagamento, ambos devem atender à interface criada.

a) Cartão de Crédito

O cartão deve conter:

Identificação do cartão

Limite total

Limite disponível

Regras:

Não permitir pagamentos acima do limite disponível

Reduzir o limite disponível após um pagamento bem-sucedido

b) PIX

O PIX deve conter:

Identificação da chave

Limite diário disponível

Regras:

Não permitir pagamentos acima do limite diário

Reduzir o limite diário após um pagamento bem-sucedido

5️⃣ Transação

Crie uma estrutura que represente uma transação financeira.

Cada transação deve conter:

Identificador

Valor

Método de pagamento utilizado

Status da operação (ex: sucesso ou falha)

6️⃣ Operação Principal

Implemente uma operação que permita ao usuário realizar um pagamento.

Essa operação deve:

Receber um método de pagamento (interface)

Receber o valor da transação

Tentar executar o pagamento

Em caso de falha, retornar erro e não registrar transação

Em caso de sucesso:

Atualizar o saldo da carteira

Registrar a transação no histórico

⚠️ Essa operação deve ser um método do usuário.

🧪 Cenários de Teste Obrigatórios

No programa principal, crie cenários que validem:

Criação de um usuário com carteira vazia

Pagamento válido com cartão

Pagamento inválido com cartão (excede limite)

Pagamento válido com PIX

Impressão do saldo final

Impressão do histórico de transações

📐 Restrições Técnicas

❌ Não utilizar herança

❌ Não utilizar switch ou if para identificar o tipo do pagamento

❌ Não utilizar bibliotecas externas

✔️ Utilizar composição

✔️ Utilizar interfaces corretamente

✔️ Utilizar tratamento de erro idiomático do Go