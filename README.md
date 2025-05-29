# 📊 Controle de Tarefas Diárias

Sistema de monitoramento e contagem de tarefas em tempo real para equipes internas de uma empresa.

---

## 🚀 Visão Geral

Esta aplicação foi desenvolvida em **Go** com foco em desempenho, organização e colaboração. Ela permite que os funcionários incrementem tarefas ao longo do dia, com **atualização em tempo real para todos os dispositivos conectados**.

Às **20:55** de cada dia, os dados são automaticamente **resetados** e um **relatório HTML** é gerado e enviado por e-mail.

---

## 🛠️ Tecnologias Utilizadas

* **Golang**
* **Gin** (framework web)
* **GORM** (ORM para banco de dados)
* **SQLite** (pode ser substituído por PostgreSQL ou outro)
* **Gorilla WebSocket** (para comunicação em tempo real)
* **Cron** com `robfig/cron` (para tarefas agendadas)
* **Go Mail** (`gomail.v2`) para envio de relatórios por e-mail

---

## 📁 Estrutura de Diretórios

```
.
├── cmd/                  # Arquivo principal (main.go)
├── internal/
│   ├── database/         # Inicialização e migração do banco
│   ├── handlers/         # Endpoints da API
│   ├── models/           # Models do GORM
│   ├── scheduler/        # Reset diário e geração de relatório
│   └── websocket/        # Lógica de WebSocket (Hub, Client, Handler)
├── templates/
│   └── report.html       # Template do relatório enviado por e-mail
└── go.mod
```

---

## 🔧 Endpoints da API

| Método | Rota                     | Descrição                               |
| ------ | ------------------------ | --------------------------------------- |
| GET    | `/employees`             | Lista todos os funcionários             |
| POST   | `/employees`             | Cadastra um novo funcionário            |
| GET    | `/tasks`                 | Lista todas as tarefas                  |
| POST   | `/tasks`                 | Cadastra uma nova tarefa                |
| POST   | `/task_counts/increment` | Incrementa uma tarefa (contagem +1)     |
| GET    | `/ws`                    | Conecta via WebSocket para atualizações |

---

## ⚡ Atualização em Tempo Real

A aplicação possui suporte a **WebSocket**. Toda vez que uma tarefa é incrementada, todos os clientes conectados recebem a nova contagem automaticamente.

### Exemplo básico de cliente WebSocket:

```html
<script>
  const ws = new WebSocket("ws://localhost:8080/ws");

  ws.onmessage = (event) => {
    const data = JSON.parse(event.data);
    console.log("Atualização recebida:", data);
  };
</script>
```

---

## ⏰ Reset Automático Diário

* O reset das contagens ocorre diariamente às **20:55**.
* Um **relatório HTML** é gerado com as contagens do dia.
* O relatório é enviado para os e-mails definidos no código.

---

## 📦 Executando o Projeto

### 1. Clone o repositório

```bash
git clone https://github.com/seu-usuario/nome-do-repositorio.git
cd nome-do-repositorio
```

### 2. Instale as dependências

```bash
go mod tidy
```

### 3. Execute a aplicação

```bash
go run cmd/main.go
```

A aplicação estará disponível em `http://localhost:8080`.

---

## ✅ Exemplos de Uso

### Incrementar contagem:

```bash
curl -X POST http://localhost:8080/task_counts/increment \
  -H "Content-Type: application/json" \
  -d '{"employee_id": 1, "task_id": 3}'
```

---

## 📬 Configuração de Email (opcional)

No agendador diário (`internal/scheduler`), configure os dados de envio:

```go
m.SetHeader("From", "seu-email@dominio.com")
m.SetHeader("To", "destino@empresa.com")
```

---

## 🤝 Contribuição

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues, propor melhorias ou enviar pull requests.

---

## 📄 Licença

Este projeto está licenciado sob a licença MIT. Veja o arquivo `LICENSE` para mais detalhes.

---

Feito com ❤️ em Go.
