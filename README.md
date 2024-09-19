# Desafio Go: Stress test

Este é um sistema desenvolvido em Go para realizar testes de carga em um serviço web. Ele permite configurar a URL do serviço, o número total de requests e o número de chamadas simultâneas. Ao final dos testes, é gerado um relatório com informações como tempo total gasto, total de requests realizados e distribuição de status HTTP.

## Pré-requisitos

Antes de começar, certifique-se de ter o seguinte instalado em sua máquina:

- Go (versão 1.22.3 ou superior)
- Docker (opcional, para execução via contêiner Docker)

## Instalação e Configuração

### Clonando o repositório

```bash
git clone https://github.com/jorgemarinho/stress-test-go.git
cd stress-test-go
```

## Compilação e Execução

### Executar diretamente do código-fonte

1. **Compilar o código-fonte:**

   ```bash
   go build -o loadtester ./cmd/loadtester
   ```

2. **Executar o teste de carga:**

   Substitua os valores de `--url`, `--requests` e `--concurrency` de acordo com suas necessidades.

   ```bash
   ./loadtester --url=http://google.com --requests=100 --concurrency=10
   ```

### Executar via Docker

1. **Construir a imagem Docker:**

   ```bash
   docker build -t stress-test-go .
   ```

2. **Executar o teste de carga via Docker:**

   Substitua os valores de `--url`, `--requests` e `--concurrency` de acordo com suas necessidades.

   ```bash
   docker run stress-test-go --url=http://google.com --requests=100 --concurrency=10
   ```

## Resultado do Teste

Após a execução, o sistema apresentará um relatório detalhado que inclui:

- Tempo total gasto na execução
- Total de requests realizados
- Quantidade de requests com status HTTP 200
- Distribuição de outros códigos de status HTTP (como 404, 500, etc.).
