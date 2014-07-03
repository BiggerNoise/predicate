@echo off
go get
go run setup.go
go test -coverprofile=coverage.out

