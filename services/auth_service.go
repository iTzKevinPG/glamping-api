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

func (as *AuthService) Authenticate(email, password string) int {
	var storedPassword string
	var userID int
	err := as.db.QueryRow("SELECT password, id FROM users WHERE email = ?", email).Scan(&storedPassword, &userID)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Usuario no encontrado")
			return 0
		}
		log.Println("Error al consultar la base de datos:", err)
		return -1
	}

	if password == storedPassword {
		return userID
	}

	return 0
}

func (as *AuthService) Registration(payload models.RegisterCredentials) int {
	var exists bool
	err := as.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = ? OR username = ?)", payload.Email, payload.FullName).Scan(&exists)

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
			log.Println("Error al insertar el usuario:", payload)
			return -1
		} else {
			return 1
		}
	}

	return 0
}
