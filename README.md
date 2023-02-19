Testar

`go test -cover -v ./src/test/... -coverpkg=./src/app/... -coverprofile=coverage.out -covermode count ./...`

Verificar cobertura

`go tool cover -func coverage.out`

Criar html da cobertura

`go tool cover -html=coverage.out -o coverage.html`

Rodar aplicacao com docker
```bash
docker compose -f docker-compose.dev.yml build --no-cache

docker compose -f docker-compose.dev.yml up
```