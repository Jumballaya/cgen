/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/wlbr/templify
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package templates

// CppMakefileTemplate is a generated function returning the template as a string.
// That string should be parsed by the functions of the golang's template package.
func CppMakefileTemplate() string {
	var tmpl = "#######\n" +
		"#\n" +
		"#   Project: {{ .Name }} {{ if .Author.Name }}\n" +
		"#   Author: {{ .Author.Name }} {{ if .Author.Email }}<{{ .Author.Email }}>{{ end }}{{ end }}\n" +
		"#\n" +
		"#######\n" +
		"\n" +
		"CXX = g++\n" +
		"CXXFLAGS = -Wall -Wextra -Werror -std=c++11 -stdlib=libc++\n" +
		"LDFLAGS = -L/usr/lib -lstdc++ -lm\n" +
		"\n" +
		"BUILD = ./bin\n" +
		"OBJ_DIR = ./obj\n" +
		"APP_DIR = $(BUILD)\n" +
		"\n" +
		"TARGET = {{ .SafeName }}\n" +
		"INCLUDE = -I./inc\n" +
		"SRC = $(wildcard src/*.cpp)\n" +
		"\n" +
		"OBJECTS = $(SRC:%.cpp=$(OBJ_DIR)/%.o)\n" +
		"DEPENDENCIES = $(OBJECTS:.o=.d)\n" +
		"\n" +
		"all: build $(APP_DIR)/$(TARGET)\n" +
		"\n" +
		"$(OBJ_DIR)/%.o: %.cpp\n" +
		"\t@mkdir -p $(@D)\n" +
		"\t$(CXX) $(CXXFLAGS) $(INCLUDE) -c $< -MMD -o $@\n" +
		"\n" +
		"$(APP_DIR)/$(TARGET): $(OBJECTS)\n" +
		"\t@mkdir -p $(@D)\n" +
		"\t$(CXX) $(CXXFLAGS) -o $(APP_DIR)/$(TARGET) $^ $(LDFLAGS)\n" +
		"\n" +
		"-include $(DEPENDENCIES)\n" +
		"\n" +
		".PHONY: all build clean debug release info\n" +
		"\n" +
		"build:\n" +
		"\t@mkdir -p $(APP_DIR)\n" +
		"\t@mkdir -p $(OBJ_DIR)\n" +
		"\n" +
		"debug: CXXFLAGS += -DDEBUG -g\n" +
		"debug: all\n" +
		"\n" +
		"release: CXXFLAGS += -O2\n" +
		"release: all\n" +
		"\n" +
		"clean:\n" +
		"\t-@rm -rvf $(OBJ_DIR)/*\n" +
		"\t-@rm -rvf $(APP_DIR)/*\n" +
		"\n" +
		"info:\n" +
		"\t@echo \"[*] Application dir: ${APP_DIR}     \"\n" +
		"\t@echo \"[*] Object dir:      ${OBJ_DIR}     \"\n" +
		"\t@echo \"[*] Sources:         ${SRC}         \"\n" +
		"\t@echo \"[*] Objects:         ${OBJECTS}     \"\n" +
		"\t@echo \"[*] Dependencies:    ${DEPENDENCIES}\""
	return tmpl
}
