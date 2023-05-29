# Soltinho JJ - Sistema de Gerenciamento de Academias de Jiu-Jitsu

O Soltinho JJ é um sistema desenvolvido em Go que permite o gerenciamento de academias de jiu-jitsu, incluindo o cadastro de academias, alunos e controle de frequência e pagamentos. O projeto utiliza a ferramenta SQLC e Docker para facilitar o gerenciamento do banco de dados.

## Funcionalidades

Criação de Usuário com Autenticação: O sistema permite a criação de usuários com diferentes funções (roles) e permissões distintas. São definidas três roles com permissões diferentes para cada tipo de usuário.

Cadastro de Academias: É possível cadastrar academias de jiu-jitsu no sistema, informando detalhes como nome, endereço e telefone. A criação de academias permite a associação de alunos a uma academia específica.

Cadastro de Alunos para Academia: O sistema permite o cadastro de alunos, associando-os a uma academia específica. As informações registradas incluem nome, idade e graduação do aluno. Isso permite uma organização adequada dos alunos em suas respectivas academias.

Autenticação JWT: Foi implementado um sistema de autenticação baseado em JWT (JSON Web Tokens). Isso garante a segurança e a autenticação dos usuários ao acessar o sistema.

Logs na Aplicação: O sistema registra logs para manter um histórico de atividades realizadas. Esses logs auxiliam na análise e solução de problemas, além de fornecerem informações sobre as ações executadas pelos usuários.

## Pré-requisitos

Os pré-requisitos mencionados são as ferramentas necessárias para o desenvolvimento do sistema Soltinho JJ. Certifique-se de ter instalado e configurado corretamente as seguintes ferramentas em seu ambiente de desenvolvimento:

Go: É a linguagem de programação utilizada no desenvolvimento do sistema. Certifique-se de ter o Go instalado em sua máquina. Você pode verificar a instalação executando o comando ``go version no terminal.``

Docker: É utilizado para facilitar o gerenciamento do banco de dados. Certifique-se de ter o Docker instalado e em execução em seu sistema. Você pode verificar a instalação executando o comando ``docker --version no terminal.``

SQLC: É uma ferramenta utilizada para simplificar o gerenciamento de consultas SQL no Go. Certifique-se de ter o SQLC instalado em sua máquina. Você pode encontrar instruções de instalação na documentação oficial do SQLC.

Framework Chi: É um framework web leve em Go utilizado para construir a API do sistema. Certifique-se de ter o framework Chi instalado. Você pode adicionar o framework Chi ao seu projeto Go executando o comando go get github.com/go-chi/chi no terminal.

JWT (JSON Web Tokens): É um método para autenticação e transmissão segura de informações entre partes por meio de tokens. Certifique-se de ter a biblioteca JWT instalada em seu projeto Go. Você pode adicionar a biblioteca JWT ao seu projeto executando o comando ``go get "github.com/go-chi/jwtauth"``
 no terminal.

Testify: É uma biblioteca em Go que fornece uma estrutura para escrever testes de unidade e asserts. Certifique-se de ter a biblioteca Testify instalada em seu projeto Go. Você pode adicionar a biblioteca Testify ao seu projeto executando o comando ``go get github.com/stretchr/testify`` no terminal.

Certifique-se de executar o comando ``go mod tidy`` após adicionar ou remover pacotes em seu projeto Go. Isso ajudará a manter o arquivo go.mod limpo e atualizado com as dependências do projeto.

Após ter todas as ferramentas e dependências instaladas, você estará pronto para prosseguir com o desenvolvimento do sistema Soltinho JJ.

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
