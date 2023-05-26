# Soltinho JJ - Sistema de Gerenciamento de Academias de Jiu-Jitsu

O Soltinho JJ é um sistema desenvolvido em Go que permite o gerenciamento de academias de jiu-jitsu, incluindo o cadastro de academias, alunos e controle de frequência e pagamentos. O projeto utiliza a ferramenta SQLC e Docker para facilitar o gerenciamento do banco de dados.

## Funcionalidades

O Soltinho JJ oferece as seguintes funcionalidades:

- Cadastro de academias de jiu-jitsu, incluindo informações como nome, endereço e telefone.
- Cadastro de alunos, com detalhes como nome, idade e graduação.
- Controle de frequência dos alunos, registrando presenças e faltas em cada aula.
- Acompanhamento dos pagamentos dos alunos, indicando se estão em dia ou com pendências.

## Pré-requisitos

Certifique-se de ter as seguintes ferramentas instaladas em seu ambiente de desenvolvimento:

- Go 
- Docker 
- SQLC 

## Configuração do ambiente

Siga os passos abaixo para configurar o ambiente de desenvolvimento:

1. Clone o repositório do projeto:

```
git clone https://github.com/seu-usuario/soltinho-jj.git
```

2. Instale as dependências do projeto:

```
go mod tidy
```

3. Execute o banco de dados PostgreSQL utilizando o Docker:

```
docker-compose up -d
```

4. Execute as migrações do banco de dados:

```
sqlc generate
```

5. Inicie o servidor da aplicação:

```
go run main.go
```

## Contribuição

Se você deseja contribuir com o projeto, fique à vontade para realizar um fork e enviar suas melhorias através de pull requests. Certifique-se de seguir as diretrizes de contribuição e adicionar testes para suas funcionalidades ou correções de bugs.
