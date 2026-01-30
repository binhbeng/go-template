package service

import (
	"github.com/binhbeng/goex/internal/model"
)

type {{.Pascal}}Service struct {
	{{.Name}}Repo *model.{{.Pascal}}Repository
}

func New{{.Pascal}}Service(
	{{.Name}}Repo *model.{{.Pascal}}Repository,
) *{{.Pascal}}Service {
	return &{{.Pascal}}Service{
		{{.Name}}Repo: {{.Name}}Repo,
	}
}