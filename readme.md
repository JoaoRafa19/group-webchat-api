# Server-sent events

Aplicação para testar o uso de Server-sent events.

## Dependencias

- Docker
- Docker-compose
- Go

## Como executar

```bash
docker-compose up -d
```

```bash
go run main.go
```

## Como testar

Abra no navegador o endereço: <http://localhost:3000>

E também abra o endereço do RabbitMQ: <http://localhost:15672>

Faça login com o usuário e senha padrão: guest/guest

Clique na aba "Queues" e depois na fila "events". Você verá as mensagens sendo enviadas para a fila.
