<br/>

<!-- <p align="center"><a href="https://newhappen.com.br" target="_blank"><img src="https://github.com/NewHappen-Company/oficial/blob/master/src/assets/logoBlue.svg?raw=true" height="70"></a></p> -->

<p align="center">
<h1 align="center">GoUniversity</h1>
</p>

<br/>

<p align="center">
    <img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go Lang" />
    <img src="https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white" alt="Postgres" />
    <img src="https://img.shields.io/badge/Docker-2CA5E0?style=for-the-badge&logo=docker&logoColor=white" alt="Docker" />
</p>

## Sobre
Uma API feita em Go sobre cadastros de matérias, servidores, departamentos e universidades

## Features
- Cadastro de universidades
- Cadastro de departamentos (em desenvolvimento)
- Cadastro de matérias (em desenvolvimento)
- Cadastro de servidores (em desenvolvimento)
- Seleção de universidades (em desenvolvimento)
- Seleção de departamentos (a fazer)
- Seleção de matérias (a fazer)
- Seleção de servidores (a fazer)

## :warning: Arquivos importantes
config.toml
-------------

O arquivo ***config.toml*** deve ser modificado com base nas configurações do seu banco de dados e porta da api.
  
Esse arquivo deve estar na pasta raiz do projeto e conter obrigatóriamente as seguintes informações:
```toml
[api]
port

[db]
host
port
user
password
database
```
>EM HIPÓTESE ALGUMA SUBA SEU ARQUIVO DE CONFIGURAÇÃO PARA O GITHUB
  
Caso seja necessário criar uma nova variável de desenvolvimento, observe que há um padrão:
```toml
[ESCOPO]
NOME_DA_VARIAVEL = "valor"
```

_Caso tenha criado uma variável, especifique isso em um Pull Request_  

-----------------

## Requisitos
- [x] Configurar config.toml
- [x] Instalar dependências com `go mod tidy`

## Rodando o backend
Execute o backend com
```shell
go run main.go
```

## Criando seu banco de dados
Para criar seu banco de dados, execute o seguinte comando:
```shell
docker run -d --name api-gouniversity -p 5433:5432 -e POSTGRES_PASSWORD=1234 postgres:13.5
```
O comando acima irá criar um container com o banco de dados postgresql na porta 5433. Para acessar o banco de dados, utilize o seguinte comando:
```shell
docker exec -it api-gouniversity psql -U postgres
```
Agora, para criar o banco de dados, execute o seguinte comando:
```sql
create database gouniversity;
```
Após a criação do banco de dados, você poderá criar um usuário para acessar o bd. Caso não queira criar um usuário, deverá utilizar o usuário padrão do postgres no seu arquivo de configuração. Para criar um usuário, execute o seguinte comando:
```sql
create user user_gouniversity;
```
Para definir uma senha para o usuário:
```sql
alter user user_gouniversity with encrypted password '1234';
```
Para dar permissões ao usuário:
```sql
grant all privileges on database gouniversity to user_gouniversity;
grant all privileges on all tables in schema public to user_gouniversity;
grant all privileges on all sequences in schema public to user_gouniversity;
```
Por fim, nos conectaremos ao banco de dados e criaremos as tabelas:
```sql
\c gouniversity
create table schools (id serial primary key, name varchar, type varchar);
```
Para sair do banco de dados, execute o comando `exit` no seu terminal.
Lembre-se de configurar o arquivo de configuração com as informações do seu banco de dados. 

## Contribuição
Fico muito feliz com a sua ajuda em meu projeto. Para que o goUniversity seja da melhor maneira possível, pedimos que leia nossa [Contribuition Guidelines](/.github/CONTRIBUITING.md).