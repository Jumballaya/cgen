#######
#
#   Project: {{ .Name }} {{ if .Author.Name }}
#   Author: {{ .Author.Name }} {{ if .Author.Email }}<{{ .Author.Email }}>{{ end }}{{ end }}
#
#######

PRODUCT := {{ .SafeName }}
BINDIR  := ./bin
INCDIR  := ./inc
SRCDIR  := ./src
OBJDIR  := ./obj

MKDIR_P = mkdir -p

CC := gcc
LINKER := gcc
INCLUDES := -I$(INCDIR) -I./
CCFLAGS := -Wall -Wextra

# Finds all .c files and puts them into SRC
SRC := $(wildcard $(SRCDIR)/*.c)
# Creates .o files for every .c file in SRC (patsubst is pattern substitution)
OBJ := $(patsubst $(SRCDIR)/%.c, $(OBJDIR)/%.o, $(SRC))
# Creates .d files (dependencies) for every .c file in SRC
DEP := $(patsubst $(SRCDIR)/%.c, $(OBJDIR)/%.d, $(SRC))

# $^ is list of dependencies and $@ is the target file
# Link all the object files
$(BINDIR)/$(PRODUCT): directories $(OBJ)
	$(LINKER) $(OBJ) -o $@

# Compile individual .c source files into object files
$(OBJDIR)/%.o: $(SRCDIR)/%.c
	$(CC) $(CCFLAGS) $(INCLUDES) -c $< -o $@


-include $(DEP)

.PHONY: directories

directories: $(OBJDIR) $(BINDIR)

$(OBJDIR):
	$(MKDIR_P) $(OBJDIR)

$(BINDIR):
	$(MKDIR_P) $(BINDIR)

.PHONY: clean

clean:
	rm -f $(OBJDIR)/*
	rm -f $(BINDIR)/$(PRODUCT)