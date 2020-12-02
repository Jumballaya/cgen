package main

import (
	"fmt"
	"os"
)

func usage() string {
	return `
Usage: cgen <Project Name> [OPTIONS]

A tool for scaffolding out C/C++ projects

Options:
    -r, --root      sets Root folder to generate files in (default: "./")
    -a, --author    sets the Author's name
    -e, --email     sets the author's Email
    -l, --lang      sets the programming language (C/C++)
	`
}

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println(usage())
		return
	}

	name := args[1]
	root := "./"
	author := ""
	email := ""
	lang := "c"

	for i, e := range args {
		if e == "-r" || e == "--root" {
			root = args[i+1]
			if root[len(root)-1] != '/' {
				root = root + "/"
			}
		}
		if e == "-a" || e == "--author" {
			author = args[i+1]
		}
		if e == "-e" || e == "--email" {
			email = args[i+1]
		}
		if e == "-l" || e == "--lang" {
			lang = args[i+1]
		}
	}

	// If root doesn't exist, create it
	if root != "./" {
		if _, err := os.Stat(root); os.IsNotExist(err) {
			os.Mkdir(root, os.ModePerm)
		}
	}

	p := NewProject(name, root, author, email, lang)
	p.Run()
}
