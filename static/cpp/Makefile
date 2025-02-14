#######
#
#   Project: {{ .Name }} {{ if .Author.Name }}
#   Author: {{ .Author.Name }} {{ if .Author.Email }}<{{ .Author.Email }}>{{ end }}{{ end }}
#
#######

CXX = g++
CXXFLAGS = -Wall -Wextra -Werror -std=c++11 -stdlib=libc++
LDFLAGS = -L/usr/lib -lstdc++ -lm

BUILD = ./bin
OBJ_DIR = ./obj
APP_DIR = $(BUILD)

TARGET = {{ .SafeName }}
INCLUDE = -I./inc
SRC = $(wildcard src/*.cpp)

OBJECTS = $(SRC:%.cpp=$(OBJ_DIR)/%.o)
DEPENDENCIES = $(OBJECTS:.o=.d)

all: build $(APP_DIR)/$(TARGET)

$(OBJ_DIR)/%.o: %.cpp
	@mkdir -p $(@D)
	$(CXX) $(CXXFLAGS) $(INCLUDE) -c $< -MMD -o $@

$(APP_DIR)/$(TARGET): $(OBJECTS)
	@mkdir -p $(@D)
	$(CXX) $(CXXFLAGS) -o $(APP_DIR)/$(TARGET) $^ $(LDFLAGS)

-include $(DEPENDENCIES)

.PHONY: all build clean debug release info

build:
	@mkdir -p $(APP_DIR)
	@mkdir -p $(OBJ_DIR)

debug: CXXFLAGS += -DDEBUG -g
debug: all

release: CXXFLAGS += -O2
release: all

clean:
	-@rm -rvf $(OBJ_DIR)/*
	-@rm -rvf $(APP_DIR)/*

info:
	@echo "[*] Application dir: ${APP_DIR}     "
	@echo "[*] Object dir:      ${OBJ_DIR}     "
	@echo "[*] Sources:         ${SRC}         "
	@echo "[*] Objects:         ${OBJECTS}     "
	@echo "[*] Dependencies:    ${DEPENDENCIES}"