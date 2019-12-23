# Cgen

Generate simple C projects with ease.

```
Usage: cgen <Project Name> [OPTIONS]

A tool for scaffolding out C projects

Options:
    -r, --root      sets Root folder to generate files in (default: "./")
    -a, --author    sets the Author's name
    -e, --email     sets the author's Email
```

use `go install` to build the binary and install it

This will create the following:

1. Makefile to build `.c` and `.h` files
2. `src/` and `inc/` folders
3. `src/main.c` with the `int main()` entry with a simple 'Hello World' program
4. `src/_.c` and `inc/_.h` files with names based off the project name
5. `.gitignore` file to ignore built binaries