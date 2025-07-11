package main

import (
	"log"
	"lucienne/config"
	"net/http"

	"lucienne/internal/handlers"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gorilla/mux"
)

const (
	MigrationsPath = "file://db/migrations"
	SeedsPath      = "file://db/seeds"

)

func main() {
	r := mux.NewRouter()

	handlers.ReturnHealth(r)
	handlers.DefineAuthors(r)

	log.Println("Rodando na porta: " + config.EnvVariables.AppPort)
	log.Fatal(http.ListenAndServe(":"+config.EnvVariables.AppPort, r))
}

func init() {
	config.EnvVariables.Load()

	m, err := migrate.New(
		MigrationsPath,
		config.EnvVariables.DatabaseURL,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Iniciando migrações...")
	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Println("Nenhuma migração pendente. Banco de dados já está atualizado.")
		} else {
			log.Fatalf("Erro ao aplicar migrações: %v", err)
		}
	} else {
		log.Println("Migrações aplicadas com sucesso.")
	}

	// Log do estado atual das migrações
	version, dirty, err := m.Version()
	if err != nil {
		log.Fatalf("Erro ao obter versão das migrações: %v", err)
	}
	log.Printf("Versão atual do banco de dados: %d, Dirty: %v", version, dirty)

	if config.EnvVariables.AppEnv == "development" {
		log.Println("Ambiente de desenvolvimento detectado. Aplicando seed...")
		seed, err := migrate.New(
			SeedsPath,
			config.EnvVariables.DatabaseURL)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Iniciando seed...")
		if err := seed.Up(); err != nil {
			if err == migrate.ErrNoChange {
				log.Println("Nenhum seed pendente. Banco de dados já está atualizado.")
			} else {
				log.Fatalf("Erro ao aplicar migrações: %v", err)
			}
		} else {
			log.Println("Seed aplicadas com sucesso.")
		}
	}

}
