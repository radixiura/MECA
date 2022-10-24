#!/bin/bash
go mod init meca

go mod tidy

cd cmd

go run main.go
