package main

import (
	"api/user-api-v1/handlers"
	"api/user-api-v1/services"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var allowedOrigins = []string{
	"http://127.0.0.1:8081",
	"http://127.0.0.1:3000",
	"http://localhost:3000",
}

func main() {
	db, err := sql.Open("mysql", "root:admin@tcp(localhost:3306)/glampingdb")
	if err != nil {
		log.Fatal("Error al abrir la base de datos:", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Error al conectar con la base de datos:", err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS migrations (
			id INT AUTO_INCREMENT,
			migration TEXT NOT NULL,
			PRIMARY KEY (id)
		)`)

	if err != nil {
		log.Fatal("Error al crear la tabla migrations:", err)
	}

	err = runMigrations(db)
	if err != nil {
		log.Fatal("Error al ejecutar migraciones:", err)
	}

	r := gin.Default()

	r.Use(corsMiddleware())
	authService := services.NewAuthService(db)
	userService := services.NewUserService(db)
	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)

	r.POST("/login", authHandler.Login)
	r.POST("/register", authHandler.Register)
	r.GET("/user", userHandler.GetUser)
	r.PUT("/user", userHandler.UpdateUser)

	r.Run("localhost:8080")
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		if origin == "" {
			c.Next()
			return
		}

		originFound := false
		for _, allowedOrigin := range allowedOrigins {
			if origin == allowedOrigin {
				originFound = true
				break
			}
		}

		if originFound {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			if c.Request.Method == "OPTIONS" {
				c.AbortWithStatus(http.StatusNoContent)
				return
			}
		} else {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Next()
	}
}

func runMigrations(db *sql.DB) error {
	migrations, err := filepath.Glob("./migrations/*.sql")
	if err != nil {
		return err
	}
	sort.Strings(migrations)

	appliedMigrations := getAppliedMigrations(db)

	for _, migration := range migrations {
		if !migrationApplied(appliedMigrations, migration) {
			err := applyMigration(db, migration)
			if err != nil {
				return err
			}
			fmt.Println("Applied migration:", migration)
		}
	}

	return nil
}

func getAppliedMigrations(db *sql.DB) []string {
	rows, err := db.Query("SELECT migration FROM migrations")
	if err != nil {
		return []string{}
	}
	defer rows.Close()

	var appliedMigrations []string
	for rows.Next() {
		var migration string
		err := rows.Scan(&migration)
		if err == nil {
			appliedMigrations = append(appliedMigrations, migration)
		}
	}
	return appliedMigrations
}

func migrationApplied(appliedMigrations []string, migrationFile string) bool {
	migrationName := filepath.Base(migrationFile)
	for _, appliedMigration := range appliedMigrations {
		if appliedMigration == migrationName {
			return true
		}
	}
	return false
}

func applyMigration(db *sql.DB, migrationFile string) error {
	migrationName := filepath.Base(migrationFile)
	migrationContent, err := os.ReadFile(migrationFile)
	if err != nil {
		return err
	}

	commands := strings.Split(string(migrationContent), ";")

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for _, cmd := range commands {
		cmd = strings.TrimSpace(cmd)
		if cmd == "" {
			continue
		}
		_, err := tx.Exec(cmd)
		if err != nil {
			return err
		}
	}

	_, err = tx.Exec("INSERT INTO migrations (migration) VALUES (?)", migrationName)
	if err != nil {
		return err
	}

	return tx.Commit()
}
