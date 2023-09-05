O arquivo `.ciignore` possui o nome dos módulos que devem ser ignorados, para saber o nome exato basta executar `go list ./...` e o nome dos módulos existentes serão exibidos

Exibir funções para criar mock das interfaces

`mockgen -source=app/controller/controller.go`
`mockgen -source=app/repository/repository.go`

Executar testes

`go test -coverprofile=coverage.out ./...`

Verificar total de cobertura

`go test -coverprofile=coverage.out ./... | go tool cover -func=coverage.out | grep total | awk '{print $3}'`

Verificar modulos sem testes

`go test -coverprofile=coverage.out ./... | grep -E "no test files" | awk '{print $2}' | grep -xvf ./app/test/.ciignore | awk '{print "NO TEST\t", $0}'`