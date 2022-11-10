#!/bin/bash

sudo systemctl start postgresql

go mod init meca

go mod tidy

cd cmd

go run main.go
