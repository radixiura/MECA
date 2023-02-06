#!/bin/bash
// Linux //
sudo apt install gcc-go
sudo apt install golang-go
sudo apt install postgres

Зайти в каталог проекта там же где этот файл

go mod init meca

go mod tidy
go mod vendor

(sql)

cd cmd && go run main.go
//* Linux *//

// Sql //
sudo systemctl start postgresql
sudo service postgres start

sudo passwd postgres -> "sir1"
su -l postgres -> "sir1"
psql

CREATE USER mecahost WITH password 'fenr1rson';
CREATE DATABASE mecadb;
\c mecadb
CREATE TABLE users( id SERIAL PRIMARY KEY, login CHARACTER VARYING(30), password CHARACTER VARYING(30), rank INTEGER );
GRANT ALL PRIVILEGES ON DATABASE mecadb to mecahost;
GRANT ALL ON users TO mecahost;

\q
exit
clear
//* Sql *//
