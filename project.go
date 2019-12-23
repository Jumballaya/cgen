package main

import (
	"os"
	"strings"
	"text/template"

	templates "github.com/jumballaya/cgen/templates/c"
)

// ProjectData is config struct for the c files
type ProjectData struct {
	Name       string
	SafeName   string
	HeaderName string
	Author     AuthorData
}

// AuthorData holds the info on the creator of the project
type AuthorData struct {
	Name  string
	Email string
}

// ProjectFile is a file's template and relative path
type ProjectFile struct {
	Template *template.Template
	Path     string
}

func newProjectFile(temp string, path string) ProjectFile {
	tmpl, err := template.New(path).Parse(temp)
	if err != nil {
		panic(err)
	}
	return ProjectFile{
		Template: tmpl,
		Path:     path,
	}
}

// Project represents a new C project
type Project struct {
	Data  ProjectData
	Files []ProjectFile
	Root  string
}

// NewProject generates a new project from a name
func NewProject(name, root, author, email string) Project {
	data := ProjectData{
		Name:       name,
		SafeName:   genSafeName(name),
		HeaderName: genHeaderName(name),
		Author:     AuthorData{author, email},
	}

	files := []ProjectFile{
		newProjectFile(templates.MakefileTemplate(), "Makefile"),
		newProjectFile(templates.MainTemplate(), "src/main.c"),
		newProjectFile(templates.ExampleHeaderTemplate(), "inc/"+data.SafeName+".h"),
		newProjectFile(templates.ExampleTemplate(), "src/"+data.SafeName+".c"),
		newProjectFile(templates.GitignoreTemplate(), ".gitignore"),
	}

	return Project{data, files, root}
}

// Run executes the templates against the project data and writes the output to files
func (p *Project) Run() {
	// Create the Directories
	for _, path := range []string{"inc", "src"} {
		os.Mkdir(p.Root+path, os.ModePerm)
	}

	// Create the Files
	for _, file := range p.Files {
		f, err := os.Create(p.Root + file.Path)
		if err != nil {
			panic(err)
		}

		err = file.Template.Execute(f, p.Data)
		if err != nil {
			panic(err)
		}
	}
}

// lowercase and camelcase .... Example Project -> example_project
func genSafeName(name string) string {
	lower := strings.ToLower(name)
	return strings.ReplaceAll(lower, " ", "_")
}

// uppercase + camelcase with a leading H_  .... Example Project -> H_EXAMPLE_PROJECT
func genHeaderName(name string) string {
	safe := genSafeName(name)
	return "H_" + strings.ToUpper(safe)
}
