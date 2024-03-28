O arquivo `.ciignore` possui o nome dos módulos que devem ser ignorados, para saber o nome exato basta executar `go list ./...` e o nome dos módulos existentes serão exibidos

## Executar testes com docker
```bash
docker compose -f docker-compose.test.yml up --build
```

## Exibir funções para criar mock das interfaces

```bash
mockgen -source=app/controller/controller.go
mockgen -source=app/repository/repository.go
```

## Executar testes

```bash
go test -coverprofile=coverage.out ./...
```

## Criar html da cobertura

```bash
go tool cover -html=coverage.out -o coverage.html
```

## Verificar cobertura (cada metodo)

```bash
go tool cover -func coverage.out
```

## Verificar cobertura (total)

```bash
go test -coverprofile=coverage.out ./... | go tool cover -func=coverage.out | grep total | awk '{print $3}'
```

## Verificar modulos sem testes

```bash
go test -coverprofile=coverage.out ./... | grep -E "no test files" | awk '{print $2}' | grep -xvf ./app/test/.ciignore | awk '{print "NO TEST\t", $0}'
```