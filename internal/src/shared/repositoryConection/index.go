// /internal/src/shared/repositoryConection/index.go
package repositoryconection

type repositoryConection interface {
	Conection()
}
type RepositoryConection struct {
	RepositoryConection repositoryConection
}

func (r *RepositoryConection) Conection() {
	r.RepositoryConection.Conection()
}
