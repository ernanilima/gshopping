Testar

`go test -cover -v ./src/test/... -coverpkg=./src/app/... -coverprofile=coverage.out -covermode count ./...`

Verificar cobertura

`go tool cover -func coverage.out`

Criar html da cobertura

`go tool cover -html=coverage.out -o coverage.html`