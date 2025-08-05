# URL Shortener

> [🇺🇸 English Version](README.md)

Um serviço de encurtamento de URLs moderno e eficiente construído em Go, utilizando Gin Framework e PostgreSQL.

## 🚀 Características

- ✅ Encurtamento de URLs com hash único
- ✅ Redirecionamento automático
- ✅ Contagem de cliques
- ✅ API RESTful
- ✅ Documentação da API com Swagger
- ✅ Middleware de logging com Zap
- ✅ Containerizado com Docker
- ✅ Banco de dados PostgreSQL
- ✅ Health check endpoint

## 🛠️ Tecnologias Utilizadas

- **Go 1.24.1** - Linguagem de programação
- **Gin** - Framework web HTTP
- **GORM** - ORM para Go
- **PostgreSQL** - Banco de dados
- **Zap** - Logger estruturado
- **Swagger** - Documentação da API
- **Docker & Docker Compose** - Containerização

## 📁 Estrutura do Projeto

```
url-shortener/
├── cmd/
│   └── shortener-api/
│       └── main.go              # Ponto de entrada da aplicação
├── docs/                        # Documentação Swagger
│   ├── docs.go                  # Documentação swagger gerada
│   ├── swagger.json             # Especificação OpenAPI JSON
│   └── swagger.yaml             # Especificação OpenAPI YAML
├── internal/
│   ├── app/
│   │   ├── model/
│   │   │   └── Url.go           # Modelo de dados da URL
│   │   └── service/
│   │       └── shortener.go     # Lógica de negócio
│   ├── server/
│   │   ├── handlers/
│   │   │   └── url.go           # Handlers HTTP
│   │   └── middleware/
│   │       └── logger.go        # Middleware de logging
│   └── storage/
│       └── postgres/
│           └── repository.go    # Repositório de dados
├── docker-compose.yml           # Orquestração de containers
├── dockerfile                   # Imagem Docker da aplicação
├── go.mod                       # Dependências do Go
└── go.sum                       # Checksums das dependências
```

## 🚀 Como Executar

### Pré-requisitos

- Go 1.24.1 ou superior
- Docker e Docker Compose
- PostgreSQL (se executar localmente)

### Executando com Docker Compose (Recomendado)

1. Clone o repositório:
```bash
git clone https://github.com/YuriGarciaRibeiro/url-shortener.git
cd url-shortener
```

2. Execute com Docker Compose:
```bash
docker-compose up --build
```

A aplicação estará disponível em `http://localhost:8080`

### 📚 Documentação da API

A documentação da API está disponível via Swagger UI em:
- **Swagger UI**: `http://localhost:8080/swagger/index.html`
- **OpenAPI JSON**: `http://localhost:8080/swagger/doc.json`
- **OpenAPI YAML**: Disponível em `docs/swagger.yaml`

### Executando Localmente

1. Clone o repositório:
```bash
git clone https://github.com/YuriGarciaRibeiro/url-shortener.git
cd url-shortener
```

2. Configure as variáveis de ambiente:
```bash
export DB_DSN="postgres://postgres:postgres@localhost:5432/url_shortener?sslmode=disable"
```

3. Instale as dependências:
```bash
go mod download
```

4. Execute a aplicação:
```bash
go run cmd/shortener-api/main.go
```

## 📋 Endpoints da API

> 💡 **Dica**: Para documentação completa da API com exemplos interativos, visite o Swagger UI em `http://localhost:8080/swagger/index.html` quando a aplicação estiver rodando.

### Health Check
```http
GET /health
```

**Resposta:**
```json
{
  "status": "ok"
}
```

### Encurtar URL
```http
POST /shorten
Content-Type: application/json

{
  "url": "https://exemplo.com/uma-url-muito-longa"
}
```

**Resposta:**
```json
{
  "hash": "abc123",
  "short_url": "http://localhost:8080/abc123"
}
```

### Redirecionamento
```http
GET /{hash}
```

Redireciona automaticamente para a URL original e incrementa o contador de cliques.

## 🗄️ Modelo de Dados

### Tabela URLs

| Campo       | Tipo      | Descrição                    |
|-------------|-----------|------------------------------|
| ID          | uint      | Chave primária               |
| Hash        | string    | Hash único da URL encurtada  |
| OriginalURL | string    | URL original                 |
| Clicks      | int       | Contador de cliques          |
| CreatedAt   | timestamp | Data de criação              |
| UpdatedAt   | timestamp | Data de atualização          |
| DeletedAt   | timestamp | Data de exclusão (soft delete) |

## 🔧 Configuração

### Variáveis de Ambiente

| Variável | Descrição | Valor Padrão |
|----------|-----------|--------------|
| `DB_DSN` | String de conexão PostgreSQL | `postgres://postgres:postgres@postgres:5432/url_shortener?sslmode=disable` |
| `PORT`   | Porta do servidor | `8080` |

## 🧪 Testando a API

### Usando cURL

**Encurtar uma URL:**
```bash
curl -X POST http://localhost:8080/shorten \
  -H "Content-Type: application/json" \
  -d '{"url": "https://github.com/YuriGarciaRibeiro/url-shortener"}'
```

**Testar redirecionamento:**
```bash
curl -I http://localhost:8080/{hash_retornado}
```

### Usando HTTPie

**Encurtar uma URL:**
```bash
http POST localhost:8080/shorten url=https://github.com/YuriGarciaRibeiro/url-shortener
```

## 🏗️ Arquitetura

O projeto segue uma arquitetura limpa com separação clara de responsabilidades:

- **cmd/**: Ponto de entrada da aplicação
- **internal/app/**: Lógica de negócio e modelos
- **internal/server/**: Camada HTTP (handlers e middleware)
- **internal/storage/**: Camada de persistência

## 🤝 Contribuindo

1. Faça um fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/MinhaFeature`)
3. Commit suas mudanças (`git commit -m 'Adiciona uma nova feature'`)
4. Push para a branch (`git push origin feature/MinhaFeature`)
5. Abra um Pull Request

## 📝 Licença

Este projeto está sob a licença MIT. Veja o arquivo `LICENSE` para mais detalhes.

## 👨‍💻 Autor

**Yuri Garcia Ribeiro**
- GitHub: [@YuriGarciaRibeiro](https://github.com/YuriGarciaRibeiro)

---

⭐ Se este projeto te ajudou, considere dar uma estrela!
