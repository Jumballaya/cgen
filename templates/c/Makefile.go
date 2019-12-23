/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/wlbr/templify
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package templates

// MakefileTemplate is a generated function returning the template as a string.
// That string should be parsed by the functions of the golang's template package.
func MakefileTemplate() string {
	var tmpl = "#######\n" +
		"#\n" +
		"#   Project: {{ .Name }} {{ if .Author.Name }}\n" +
		"#   Author: {{ .Author.Name }} {{ if .Author.Email }}<{{ .Author.Email }}>{{ end }}{{ end }}\n" +
		"#\n" +
		"#######\n" +
		"\n" +
		"PRODUCT := {{ .SafeName }}\n" +
		"BINDIR  := ./bin\n" +
		"INCDIR  := ./inc\n" +
		"SRCDIR  := ./src\n" +
		"OBJDIR  := ./obj\n" +
		"\n" +
		"MKDIR_P = mkdir -p\n" +
		"\n" +
		"CC := gcc\n" +
		"LINKER := gcc\n" +
		"INCLUDES := -I$(INCDIR) -I./\n" +
		"CCFLAGS := -Wall -Wextra\n" +
		"\n" +
		"# Finds all .c files and puts them into SRC\n" +
		"SRC := $(wildcard $(SRCDIR)/*.c)\n" +
		"# Creates .o files for every .c file in SRC (patsubst is pattern substitution)\n" +
		"OBJ := $(patsubst $(SRCDIR)/%.c, $(OBJDIR)/%.o, $(SRC))\n" +
		"# Creates .d files (dependencies) for every .c file in SRC\n" +
		"DEP := $(patsubst $(SRCDIR)/%.c, $(OBJDIR)/%.d, $(SRC))\n" +
		"\n" +
		"# $^ is list of dependencies and $@ is the target file\n" +
		"# Link all the object files\n" +
		"$(BINDIR)/$(PRODUCT): directories $(OBJ)\n" +
		"\t$(LINKER) $(OBJ) -o $@\n" +
		"\n" +
		"# Compile individual .c source files into object files\n" +
		"$(OBJDIR)/%.o: $(SRCDIR)/%.c\n" +
		"\t$(CC) $(CCFLAGS) $(INCLUDES) -c $< -o $@\n" +
		"\n" +
		"\n" +
		"-include $(DEP)\n" +
		"\n" +
		".PHONY: directories\n" +
		"\n" +
		"directories: $(OBJDIR) $(BINDIR)\n" +
		"\n" +
		"$(OBJDIR):\n" +
		"\t$(MKDIR_P) $(OBJDIR)\n" +
		"\n" +
		"$(BINDIR):\n" +
		"\t$(MKDIR_P) $(BINDIR)\n" +
		"\n" +
		".PHONY: clean\n" +
		"\n" +
		"clean:\n" +
		"\trm -f $(OBJDIR)/*\n" +
		"\trm -f $(BINDIR)/$(PRODUCT)"
	return tmpl
}
