#!/usr/bin/env bash

go mod init github.com/jestebanrods/edu-golang/experimental
go mod vendor
go list -m all
go list -m -u all
go install github.com/stretchr/testify/mock
go get github.com/stretchr/testify@v1.3.0
