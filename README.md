# http-server-projeto-korp

Serviço HTTP em Go com infraestrutura containerizada, monitoramento e automação completa via Ansible.

## Tecnologias

- **Go 1.23** — servidor HTTP
- **Docker + Docker Compose** — containerização
- **NGINX** — proxy reverso
- **Prometheus** — coleta de métricas
- **Grafana** — visualização de métricas
- **Ansible** — automação do provisionamento

## Arquitetura
Internet

│

▼

[NGINX :80]  ──proxy──▶  [Go Service :8080]

│

/metrics

│

[Prometheus :9090]

│

[Grafana :3000]

Todos os serviços comunicam-se via rede Docker bridge `korp-network`. O serviço Go não expõe portas ao host — só é acessível via NGINX.

## Estrutura do Projeto
http-server-projeto-korp/

├── app/                        # Código-fonte Go

│   ├── main.go

│   ├── go.mod

│   └── go.sum

├── ansible/                    # Automação

│   ├── playbook.yml

│   └── inventory.ini

├── grafana/provisioning/       # Provisionamento automático Grafana

│   ├── datasources/

│   └── dashboards/

├── nginx/conf.d/               # Configuração proxy reverso

├── prometheus/                 # Configuração de scrape

├── docker-compose.yml

├── Dockerfile

└── README.md

## Pré-requisitos

- Linux (Ubuntu 24.04 recomendado)
- Ansible Core >= 2.16
- Coleção Ansible: `ansible-galaxy collection install community.docker`
- Git

> O Docker será instalado automaticamente pelo playbook caso não esteja presente.

## Provisionamento com único comando

```bash
ansible-playbook -i ansible/inventory.ini ansible/playbook.yml --ask-become-pass
```

O playbook executa automaticamente:

1. Instalação do Docker Engine e plugin Compose
2. Build da imagem do serviço Go
3. Criação da rede bridge e subida de todos os containers
4. Configuração do NGINX como proxy reverso
5. Configuração do Prometheus e Grafana
6. Validação via requisição HTTP com exibição da resposta no console

## Endpoints

| Serviço | URL | Descrição |
|---|---|---|
| API (via NGINX) | `http://localhost:80/projeto-korp` | Endpoint principal |
| Métricas | `http://localhost:9090` | Prometheus |
| Dashboard | `http://localhost:3000` | Grafana (admin/admin) |

## Validação manual

```bash
curl http://localhost:80/projeto-korp
```

Resposta esperada:
```json
{
  "nome": "Projeto Korp",
  "horario": "2026-06-15T20:45:38Z"
}
```

## Métricas implementadas

| Métrica | Tipo | Descrição |
|---|---|---|
| `service_up` | Gauge | Disponibilidade do serviço (1=online, 0=offline) |
| `http_requests_total` | Counter | Volume de requisições por endpoint, método e status |

## Dashboard Grafana

O dashboard é provisionado automaticamente via arquivos em `grafana/provisioning/`.
Acesse `http://localhost:3000` com `admin/admin` e navegue em **Dashboards → Korp → http-server-projeto-korp**.

## Derrubar o ambiente

```bash
docker compose down
```