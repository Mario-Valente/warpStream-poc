# warpStream-poc

Este repositório é uma prova de conceito (PoC) para integração e uso do WarpStream como plataforma de streaming de dados, utilizando infraestrutura Kubernetes (Kind), Helm, Terraform e aplicações Go para produção e consumo de mensagens Kafka-like.

## Estrutura do Projeto

```
.
├── apps/
│   ├── consumer/      # Aplicação Go que consome mensagens do WarpStream
│   ├── producer/      # Aplicação Go que produz mensagens para o WarpStream
│   └── k8s/           # Manifestos Kubernetes para deploy das aplicações
├── infra/
│   ├── agents/        # Helm chart e valores para deploy do WarpStream Agent
│   ├── kind/          # Configuração do cluster Kind
│   └── terraform/     # Scripts Terraform para provisionamento de recursos AWS (S3)
├── scripts/           # Scripts de automação para setup e deploy
├── docker-compose.yaml
├── makefile
└── README.md
```

## Pré-requisitos

- [Docker](https://www.docker.com/)
- [Kind](https://kind.sigs.k8s.io/)
- [kubectl](https://kubernetes.io/docs/tasks/tools/)
- [Helm](https://helm.sh/)
- [Terraform](https://www.terraform.io/)
- [Go](https://golang.org/) >= 1.24

## Setup do Ambiente

1. **Provisionar recursos AWS (opcional):**

   Configure suas credenciais AWS e execute:

   ```sh
   cd infra/terraform
   terraform init
   terraform apply
   ```

2. **Criar cluster Kubernetes local e instalar WarpStream Agent:**

   Use o makefile para automatizar o setup:

   ```sh
   make install
   ```

   Isso irá:
   - Criar um cluster Kind usando a configuração em kind-config.yaml
   - Instalar o WarpStream Agent via Helm com os valores de values.yaml

3. **Deploy das aplicações Go (producer/consumer):**

   ```sh
   make deploy
   ```

   Isso aplica os manifestos em manifest.yaml.

## Aplicações

- **Producer:** Envia mensagens aleatórias para o tópico configurado.
- **Consumer:** Consome e imprime mensagens do tópico.

Ambas as aplicações usam variáveis de ambiente para configuração do tópico e broker, conforme definido nos manifestos Kubernetes.

## Variáveis de Ambiente

- `TOPIC_NAME`: Nome do tópico Kafka/WarpStream.
- `BROKER_URL`: URL do broker (ex: `warpstream-agent:9092`).

## Observações

- O bucket S3 e as credenciais AWS devem ser válidos para o funcionamento do WarpStream Agent.
- Os scripts e manifestos estão preparados para ambiente de desenvolvimento/local, podendo ser adaptados para produção.

## Referências

- [WarpStream Docs](https://docs.warpstream.com/)
- [Kind](https://kind.sigs.k8s.io/)
- [Helm](https://helm.sh/)
- [Terraform AWS Provider](https://registry.terraform.io/providers/hashicorp/aws/latest/docs)
