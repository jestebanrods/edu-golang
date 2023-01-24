#!/usr/bin/env bash

# github.com/husobee/vestigo
# github.com/stretchr/testify/mock

go test -v -cover

go test -coverprofile=c.out
go tool cover -html=c.out

godoc -http ':8080'
