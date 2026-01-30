package handler

import (
	"github.com/binhbeng/goex/internal/service"
)

type {{.Pascal}}Handler struct {
	{{.Name}}Service *service.{{.Pascal}}Service
}

func New{{.Pascal}}Handler({{.Name}}Service *service.{{.Pascal}}Service) *{{.Pascal}}Handler {
	return &{{.Pascal}}Handler{
		{{.Name}}Service: {{.Name}}Service,
	}
}
