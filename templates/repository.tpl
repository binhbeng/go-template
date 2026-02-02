package repository

type {{.Pascal}}Repository struct {
	*Repository
}

func New{{.Pascal}}Repository(r *Repository) *{{.Pascal}}Repository {
	return &{{.Pascal}}Repository{
		Repository: r,
	}
}

func (m *{{.Pascal}}Repository) TableName() string {
	return "{{.Name}}s"
}