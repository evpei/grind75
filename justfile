[private]
default:
  @just --list

test: 
  go test ./...

alias fmt := format

format: 
  go fmt ./...
