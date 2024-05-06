package services

import (
	"api/user-api-v1/models"
	"database/sql"
	"log"
)

type AuthService struct {
	db *sql.DB
}

func NewAuthService(db *sql.DB) *AuthService {
	return &AuthService{
		db: db,
	}
}

func (as *AuthService) Authenticate(email, password string) bool {
	var storedPassword string
	err := as.db.QueryRow("SELECT password FROM users WHERE email = ?", email).Scan(&storedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Usuario no encontrado")
			return false
		}
		log.Println("Error al consultar la base de datos:", err)
		return false
	}

	return password == storedPassword
}

func (as *AuthService) Registration(payload models.RegisterCredentials) int {
	var exists bool
	err := as.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)", payload.Email).Scan(&exists)

	if err != nil {
		log.Println("Error al consultar la base de datos:", err)
		return -1
	}

	if !exists {
		_, errI := as.db.Exec("INSERT INTO users (username, email, address, phone, password) VALUES (?, ?, ?, ?, ?)",
			payload.FullName,
			payload.Email,
			payload.Address,
			payload.Phone,
			payload.Password)

		if errI != nil {
			log.Println("Error al insertar el usuario:", err)
			return -1
		} else {
			return 1
		}
	}

	return 0
}
