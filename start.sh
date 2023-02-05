#!/bin/bash
sudo apt install gcc-go
sudo apt install golang-go

sudo systemctl start postgresql

Зайти в каталог проекта там же где этот файл

go mod init meca

go mod tidy
go mod vendor

cd cmd && go run main.go
