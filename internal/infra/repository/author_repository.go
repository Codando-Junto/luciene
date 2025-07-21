package repository

import (
	"context"
	"lucienne/internal/domain"
	"lucienne/internal/infra/database"
)

const (
	// Não precisamos retornar o ID por enquanto, então usamos um INSERT simples.
	createAuthorQuery = `INSERT INTO authors (name) VALUES ($1)`
)

// AuthorRepository define a interface para as operações de autor no banco de dados.
type AuthorRepository interface {
	CreateAuthor(ctx context.Context, author *domain.Author) error
}

// PostgresAuthorRepository é a implementação do AuthorRepository para o PostgreSQL.
type PostgresAuthorRepository struct {
	// No futuro, podemos adicionar o pool de conexões aqui.
}

// NewPostgresAuthorRepository cria uma nova instância do repositório.
func NewPostgresAuthorRepository() *PostgresAuthorRepository {
	return &PostgresAuthorRepository{}
}

// CreateAuthor insere um novo autor no banco de dados.
func (r *PostgresAuthorRepository) CreateAuthor(ctx context.Context, author *domain.Author) error {
	_, err := database.Conn.Exec(ctx, createAuthorQuery, author.Name)
	if err != nil {
		return err
	}
	return nil
}
