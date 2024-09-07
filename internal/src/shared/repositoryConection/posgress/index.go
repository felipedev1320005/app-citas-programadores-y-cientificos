// /internal/src/shared/repositoryConection/posgress/index.go
package posgress

import (
	"fmt"
	"go-rest/internal/src/users/domain"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PosgressRepositoryConection struct{}

var Db *gorm.DB

func (p *PosgressRepositoryConection) Conection() {
	// Configuración de la base de datos PostgreSQL con Gorm
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	var err error
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
	}
	log.Println("Auto migración de la base de datos")
	Db.AutoMigrate(&domain.User{}, &domain.Profile{}, &domain.Preferences{})
}
