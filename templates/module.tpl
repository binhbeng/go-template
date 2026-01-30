package app

import (
	"github.com/binhbeng/goex/data"
	"github.com/binhbeng/goex/internal/handler"
	"github.com/binhbeng/goex/internal/model"
	"github.com/binhbeng/goex/internal/service"
)

type {{.Pascal}}Module struct {
	{{.Name}}Handler *handler.{{.Pascal}}Handler
}

func New{{.Pascal}}Module() *{{.Pascal}}Module {
	repository := model.NewRepository(data.PostgreDB)
	{{.Name}}Repository := model.New{{.Pascal}}Repository(repository)
	{{.Name}}Service := service.New{{.Pascal}}Service({{.Name}}Repository)
	{{.Name}}Handler := handler.New{{.Pascal}}Handler({{.Name}}Service)
	return &{{.Pascal}}Module{
        {{.Name}}Handler: {{.Name}}Handler,
    }
}

func (m *{{.Pascal}}Module) Handler() *handler.{{.Pascal}}Handler {
	return m.{{.Name}}Handler
}