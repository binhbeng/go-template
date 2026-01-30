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
		// cleanModule(name)
	default:
		fmt.Println("Unknown command:", cmd)
	}
}

func generateModule(name string) {
	templates := []string{"dto.tpl", "handler.tpl", "service.tpl", "model.tpl", "module.tpl"}

	for _, tpl := range templates {
		var dstDir string
		switch tpl {
		case "dto.tpl":
			dstDir = filepath.Join("internal", "dto")
		case "handler.tpl":
			dstDir = filepath.Join("internal", "handler")
		case "service.tpl":
			dstDir = filepath.Join("internal", "service")
		case "model.tpl":
			dstDir = filepath.Join("internal", "model")
		case "module.tpl":
			dstDir = filepath.Join("internal", "app")
		}
		os.MkdirAll(dstDir, 0755)

		srcPath := filepath.Join("templates/", tpl)
		var dstFile string
		if tpl == "module.tpl" {
			dstFile = name + "_module.go"
		} else {
			dstFile = name + ".go"
		}

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
		filepath.Join("internal", "dto", name+".go"),
		filepath.Join("internal", "handler", name+".go"),
		filepath.Join("internal", "service", name+".go"),
		filepath.Join("internal", "model", name+".go"),
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
