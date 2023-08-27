Verificar modulos sem testes

`go test -coverprofile=coverage.out ./... | grep -E "no test files" | awk '{print $2}' | grep -xvf ./app/test/.ciignore | awk '{print "NO TEST\t", $0}'`

O arquivo `.ciignore` possui o nome dos módulos que devem ser ignorados, para saber o nome exato basta executar `go list ./...` e o nome dos módulos existentes serão exibidos

Exibir funções mocar interface

`mockgen -source=app/controller/controller.go`
`mockgen -source=app/repository/repository.go`