# GShopping

Este é um projeto em que desenvolvi uma API utilizando a linguagem Go. O principal propósito da aplicação é armazenar informações sobre produtos, incluindo seus códigos de barras e a respectiva marca a que pertencem. A API possibilita a fácil pesquisa e recuperação desses dados.

#### WEB

A parte **[WEB](https://github.com/ernanilima/wshopping)** do projeto oferece ao administrador a opção de gerenciar todos os cadastros

#### Mobile

A parte **Mobile** (_em desenvolvimento_) do projeto oferece a opção de pesquisa por código de barras, os usuários podem facilmente encontrar e registrar produtos utilizando seus dispositivos móveis.

## Tecnologias utilizadas

- [Go 1.20](https://go.dev/doc/go1.20) (Golang)
- [Docker](https://hub.docker.com/r/ernanilima/gshopping/tags)
- [Goose](https://github.com/pressly/goose)
- [Chi](https://github.com/go-chi/chi)
- [Testify](https://github.com/stretchr/testify)
- [Testcontainer](https://golang.testcontainers.org/quickstart)
- [Mockgen](https://github.com/uber-go/mock)
- [Postgres](https://www.postgresql.org)

## Rodar aplicacao com docker

```bash
docker compose -f docker-compose.dev.yml up --build

OU

docker compose -f docker-compose.dev.yml build --no-cache
docker compose -f docker-compose.dev.yml up
```

## Autor

Ernani Lima - https://ernanilima.com.br

Este projeto integra meu portfólio pessoal, e seria um prazer receber seu feedback em relação ao projeto, código, estrutura ou qualquer observação construtiva que contribua para o meu aprimoramento como desenvolvedor!

Fique à vontade para entrar em contato por meio do [LinkedIn](https://www.linkedin.com/in/ernanilima).

Além disso, este projeto está disponível para seu uso da maneira que melhor lhe convier, seja para fins de estudo ou para implementar melhorias!

Aproveite!

## Testes

Para saber mais sobre os testes, basta clicar [AQUI](https://github.com/ernanilima/gshopping/tree/main/app/test)
