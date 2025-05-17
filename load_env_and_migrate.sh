#!/bin/bash
# Script ini untuk load environment dari .env lalu menjalankan aplikasi Go Anda

if [ -f .env ]; then
  export $(grep -v '^#' .env | xargs)
fi

go run etc/migrate/migrate.go
