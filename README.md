Arquitetura Hexagonal

Projeto utilizado para prática de uma arquitetura hexagonal, na linguagem GO, utilizando
funções para criação, pesquisa, habilitar e desabilitar um produto (informar name e price para criação e id para demais operações). \
Criado adapters para banco de dados, CLI e Web Server \
Nesse projeto se utiliza a biblioteca cobra para rodar comando via CLI \
Exemplos de comandos: \
    - Para subir web server: go run main.go http (pode utilizar postman para executar as 4 funções em produtos) \
    - Para criar um produto via linha de comando: go run main.go cli -a=create -n="Product CLI" -p=25.0

Setup:
docker compose up

Run Tests:
go test ou go test ./... (se dentro do container)
