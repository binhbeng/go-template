package service

import (
	"github.com/binhbeng/goex/internal/model/repository"
)

type {{.Pascal}}Service struct {
	{{.Name}}Repo *repository.{{.Pascal}}Repository
}

func New{{.Pascal}}Service(
	{{.Name}}Repo *repository.{{.Pascal}}Repository,
) *{{.Pascal}}Service {
	return &{{.Pascal}}Service{
		{{.Name}}Repo: {{.Name}}Repo,
	}
}