package main

//go:generate templify -p templates -o ./templates/c/Main.go static/c/src/main.c
//go:generate templify -p templates -o ./templates/c/Example.go static/c/src/example.c
//go:generate templify -p templates -o ./templates/c/ExampleHeader.go static/c/inc/example.h
//go:generate templify -p templates -o ./templates/c/Gitignore.go static/c/.gitignore
//go:generate templify -p templates -o ./templates/c/Makefile.go static/c/Makefile
