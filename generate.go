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
		fmt.Println("Usage: go run generate.go module <name>")
		return
	}

	cmd := os.Args[1]
	name := strings.ToLower(os.Args[2])

	switch cmd {
	case "module":
		generateModule(name)
	default:
		fmt.Println("Unknown command:", cmd)
	}
}

func generateModule(name string) {
	dstDir := filepath.Join("modules", name)
	os.MkdirAll(dstDir, 0755)

	templates := []string{"handler.tpl", "service.tpl", "model.tpl"}

	for _, tpl := range templates {
		srcPath := filepath.Join("templates/", tpl)
		dstFile := strings.Replace(tpl, ".tpl", ".go", 1)
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
	}

	fmt.Println("Created module:", name)
}

func main() {
	run()
}
