package main

//go:generate templify -p templates -o ./templates/CMain.go static/c/src/main.c
//go:generate templify -p templates -o ./templates/CExample.go static/c/src/example.c
//go:generate templify -p templates -o ./templates/CExampleHeader.go static/c/inc/example.h
//go:generate templify -p templates -o ./templates/CGitignore.go static/c/.gitignore
//go:generate templify -p templates -o ./templates/CMakefile.go static/c/Makefile

//go:generate templify -p templates -o ./templates/CppMakefile.go static/cpp/Makefile
//go:generate templify -p templates -o ./templates/CppVectorHeader.go static/cpp/inc/Vector.h
//go:generate templify -p templates -o ./templates/CppMain.go static/cpp/src/main.cpp
//go:generate templify -p templates -o ./templates/CppGitignore.go static/cpp/.gitignore
