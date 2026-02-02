//go:build ignore
// +build ignore

//go:generate go run generate.go module order

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Data struct {
	Name   string
	Pascal string
}

func toPascal(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}

func run() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run generate.go <command> <name>")
		fmt.Println("Commands: module, clean")
		return
	}

	cmd := os.Args[1]
	name := strings.ToLower(os.Args[2])

	switch cmd {
	case "module":
		generateModule(name)
	case "clean":
		fmt.Println("Maintenance:.....", name)
		cleanModule(name)
	default:
		fmt.Println("Unknown command:", cmd)
	}
}

func generateModule(name string) {
	templates := []string{"dto.tpl", "handler.tpl", "service.tpl", "entity.tpl", "repository.tpl", "module.tpl"}
	var dstFile string

	for _, tpl := range templates {
		var dstDir string
		switch tpl {
		case "dto.tpl":
			dstDir = filepath.Join("internal", "dto")
			dstFile = name + "_dto.go"
		case "handler.tpl":
			dstDir = filepath.Join("internal", "handler")
			dstFile = name + "_handler.go"
		case "service.tpl":
			dstDir = filepath.Join("internal", "service")
			dstFile = name + "_service.go"
		case "entity.tpl":
			dstDir = filepath.Join("internal", "model", "entity")
			dstFile = name + "_entity.go"
		case "repository.tpl":
			dstDir = filepath.Join("internal", "model", "repository")
			dstFile = name + "_repository.go"
		case "module.tpl":
			dstDir = filepath.Join("internal", "app")
			dstFile = name + "_module.go"
		}
		os.MkdirAll(dstDir, 0755)

		srcPath := filepath.Join("templates/", tpl)

		dstPath := filepath.Join(dstDir, dstFile)

		t, err := template.ParseFiles(srcPath)
		if err != nil {
			panic(err)
		}

		f, err := os.Create(dstPath)
		if err != nil {
			panic(err)
		}

		t.Execute(f, Data{
			Name:   name,
			Pascal: toPascal(name),
		})

		f.Close()

		fmt.Printf("Created file: %s\n", dstPath)
	}

	fmt.Println("Generated files for module:", name)
}

func cleanModule(name string) {
	files := []string{
		filepath.Join("internal", "dto", name+"_dto.go"),
		filepath.Join("internal", "handler", name+"_handler.go"),
		filepath.Join("internal", "service", name+"_service.go"),
		filepath.Join("internal", "model", "entity", name+"_entity.go"),
		filepath.Join("internal", "model", "repository", name+"_repository.go"),
		filepath.Join("internal", "app", name+"_module.go"),
	}

	for _, file := range files {
		if err := os.Remove(file); err != nil && !os.IsNotExist(err) {
			fmt.Printf("Error removing %s: %v\n", file, err)
		} else if err == nil {
			fmt.Printf("Removed %s\n", file)
		}
	}

	fmt.Println("Cleaned files for module:", name)
}

func main() {
	run()
}
