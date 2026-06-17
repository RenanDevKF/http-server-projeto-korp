# HTTP Server Projeto Korp

![Status](https://img.shields.io/badge/status-concluído-brightgreen)
![Go](https://img.shields.io/badge/Go-1.23-00ADD8)
![Docker](https://img.shields.io/badge/Docker-enabled-2496ED)
![Prometheus](https://img.shields.io/badge/Prometheus-enabled-E6522C)
![Grafana](https://img.shields.io/badge/Grafana-enabled-F46800)
![Ansible](https://img.shields.io/badge/Ansible-automated-EE0000)

---

# Objetivo

Projeto desenvolvido como solução para o desafio técnico da vaga de Estágio DevOps da Korp.

A solução consiste em um serviço HTTP desenvolvido em Go, containerizado com Docker, publicado através de um proxy reverso NGINX, monitorado com Prometheus e Grafana, e provisionado automaticamente utilizando Ansible.

---

# Requisitos Atendidos

* [x] Serviço HTTP em Go
* [x] Endpoint GET `/projeto-korp`
* [x] Resposta JSON contendo nome e horário UTC
* [x] Containerização com Docker
* [x] Proxy reverso com NGINX
* [x] Rede Docker Bridge
* [x] Monitoramento com Prometheus
* [x] Dashboard com Grafana
* [x] Automação com Ansible
* [x] Provisionamento completo através de um único comando

---

# Pré-requisitos

* Linux (Ubuntu 24.04 LTS recomendado) ou WSL2
* Git
* Ansible Core >= 2.16
* Acesso à internet para download das dependências

> O Docker não precisa estar previamente instalado. O playbook realiza sua instalação automaticamente.

---

# Arquitetura da Solução

```text
Cliente
   │
   ▼
NGINX (porta 80)
   │
   ▼
HTTP Server Go (porta 8080)
   │
   ├── /projeto-korp → Retorna JSON
   └── /metrics      → Expõe métricas
            │
            ▼
      Prometheus (9090)
      Coleta métricas
            │
            ▼
       Grafana (3000)
      Visualiza métricas
```

Todos os serviços comunicam-se através da rede Docker Bridge `korp-network`.

O serviço Go não expõe portas diretamente ao host, sendo acessível apenas através do NGINX, que atua como proxy reverso.

---

# Tecnologias Utilizadas

## Linguagem

* Go 1.23

## Containers e Infraestrutura

* Docker
* Docker Compose
* NGINX

## Observabilidade

* Prometheus
* Grafana

## Automação

* Ansible

## Sistema Operacional

* Ubuntu 24.04 LTS
* Compatível com WSL2

---

# Estrutura do Projeto

```text
.
├── app
│   ├── main.go
│   ├── go.mod
│   └── go.sum
│
├── nginx
│   └── conf.d
│       └── http-server-projeto-korp.conf
│
├── prometheus
│   └── prometheus.yml
│
├── grafana
│   └── provisioning
│       ├── datasources
│       │   └── datasources.yml
│       └── dashboards
│           ├── dashboards.yml
│           └── http-server-projeto-korp-dashboard.json
│
├── ansible
│   ├── inventory.ini
│   └── playbook.yml
│
├── Dockerfile
├── docker-compose.yml
└── README.md
```

---

# Endpoints

| Serviço         | URL                             | Descrição                     |
| --------------- | ------------------------------- | ----------------------------- |
| API (via NGINX) | `http://localhost/projeto-korp` | Endpoint principal            |
| Prometheus      | `http://localhost:9090`         | Coleta e consulta de métricas |
| Grafana         | `http://localhost:3000`         | Dashboard de observabilidade  |

---

# Endpoint da Aplicação

## GET /projeto-korp

Resposta esperada:

```json
{
  "nome": "Projeto Korp",
  "horario": "2026-06-16T23:24:31Z"
}
```

O campo `horario` é gerado dinamicamente em UTC utilizando o padrão RFC3339.

---

# Métricas Implementadas

| Métrica               | Tipo    | Descrição                                                         |
| --------------------- | ------- | ----------------------------------------------------------------- |
| `service_up`          | Gauge   | Disponibilidade do serviço (1 = online, 0 = offline)              |
| `http_requests_total` | Counter | Volume de requisições agrupado por endpoint, método e status HTTP |

As métricas são expostas através do endpoint:

```text
http://localhost:8080/metrics
```

> Acessível diretamente apenas dentro da rede Docker. O Prometheus realiza o scrape internamente via `http://http-server-projeto-korp:8080/metrics`.

---

# Observabilidade

## Prometheus

Responsável pela coleta periódica das métricas expostas pela aplicação.

Acesso:

```text
http://localhost:9090
```

## Grafana

Responsável pela visualização das métricas coletadas.

Acesso:

```text
http://localhost:3000
```

Credenciais padrão: `admin` / `admin`

Dashboard provisionado automaticamente contendo:

* Disponibilidade do serviço
* Total acumulado de requisições
* Taxa de requisições por segundo

---

# Execução com Docker Compose

Subir todo o ambiente:

```bash
docker compose up --build -d
```

Verificar containers:

```bash
docker compose ps
```

Parar ambiente:

```bash
docker compose down
```

---

# Execução com Ansible

Instalar a coleção necessária:

```bash
ansible-galaxy collection install community.docker
```

Provisionar todo o ambiente automaticamente:

```bash
ansible-playbook -i ansible/inventory.ini ansible/playbook.yml --ask-become-pass
```

O playbook executa automaticamente:

1. Instalação do Docker Engine e Docker Compose
2. Build da imagem do serviço Go
3. Criação da rede Bridge e inicialização dos containers
4. Configuração do NGINX como proxy reverso
5. Configuração do Prometheus e Grafana
6. Validação do endpoint HTTP
7. Exibição da resposta da aplicação no console

---

# Validação

### Aplicação

```bash
curl http://localhost/projeto-korp
```

### Prometheus

```bash
curl http://localhost:9090/-/ready
```

### Grafana

```text
http://localhost:3000
```

---

# Principais Decisões Técnicas

## Multi-stage Docker Build

Utilizado para separar o ambiente de build do ambiente de execução, reduzindo o tamanho da imagem final e melhorando a segurança.

## Container executando como usuário não-root

Implementado para reduzir riscos de segurança e seguir boas práticas de execução em containers.

## NGINX como Proxy Reverso

Utilizado para desacoplar o acesso externo da aplicação e simular uma arquitetura próxima de ambientes produtivos.

## Prometheus + Grafana

Implementados para prover observabilidade da aplicação, atendendo aos requisitos de monitoramento de disponibilidade e volume de requisições.

## Ansible

Responsável pela automação completa do provisionamento da infraestrutura através de um único comando.

---

# Autor

**Renan Kirchmaier Fayer**

GitHub: https://github.com/RenanDevKF