# tutorial-goapi

Projeto do tutorial de API com Golang.

# Stack

* Echo Framework
* Gmock
* Gomega
* PostgreSQL

## Configurando Banco

Edite o arquivo .env para ajustar os parametros de conexao do banco de dados.

## Baixando as dependencias

```shell
$ go mod download 
```

## Executar em modo Desenvolvimento

```shell
$ go run cmd/main.go
```

## Executando os testes

```
$ ./scripts/cover-test.sh
```

## Gerando Binario

```shell
$ ./scripts/build.sh
```
