#!/bin/sh

docker-compose up -d

# dep ensure -update -v
dep ensure

ENV=dev go run main.go
