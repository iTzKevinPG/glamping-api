package services

import (
	"api/user-api-v1/models"
	"database/sql"
	"log"
)

type UserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		db: db,
	}
}

func (as *UserService) GetUserById(id int) *models.UserResponse {
	var user models.UserResponse
	err := as.db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(
		&user.Id, &user.FullName, &user.Email, &user.Address, &user.Phone, &user.Password, &user.PayMethodId,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Usuario no encontrado")
			return nil
		}
		log.Println("Error al consultar la base de datos:", err)
		return nil
	}

	return &user
}

func (as *UserService) UpdateUserById(user models.UpdateUser) int {
	stmt, err := as.db.Prepare("UPDATE users SET username=?, email=?, address=?, phone=?, password=?, payMethodId=? WHERE id=?")
	if err != nil {
		log.Println("Error preparing update statement:", err)
		return -1
	}

	result, err := stmt.Exec(user.FullName, user.Email, user.Address, user.Phone, user.Password, user.PayMethodId, user.Id)
	if err != nil {
		log.Println("Error updating user:", err)
		return -1
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error getting rows affected:", err)
		return -1
	}

	if rowsAffected == 0 {
		log.Println("User not found")
		return 0
	}

	return 1
}
