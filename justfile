[private]
default:
  @just --list

test: 
  go test ./...

alias fmt := format

format: 
  go fmt ./...

[linux]
@problems:
    xdg-open "https://www.techinterviewhandbook.org/grind75" 
