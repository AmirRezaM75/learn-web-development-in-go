package controllers

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"net/http"
)

type UserController struct{}

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	SSLMode  bool
}

type User struct {
	id       int
	name     string
	email    string
	password string
}

func (c PostgresConfig) toString() string {
	SSLMode := "disable"
	if c.SSLMode {
		SSLMode = "enable"
	}
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.Database, SSLMode)
}

func (uc UserController) Store(w http.ResponseWriter, r *http.Request) {
	db := getDatabase()

	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(128) NOT NULL,
			email TEXT UNIQUE NOT NULL,
		    password TEXT NOT NULL
		)
	`)

	if err != nil {
		panic(err)
	}

	email := r.FormValue("email")
	name := r.FormValue("name")
	password := r.FormValue("password")

	row := db.QueryRow(`
		INSERT INTO users (email, name, password)
		VALUES ($1, $2, $3)
		RETURNING id
	`, email, name, password)

	//  If this error is not nil, this error will also be returned from Scan.
	row.Err()
	var id int
	err = row.Scan(&id)

	if err != nil {
		panic(err)
	}

	row = db.QueryRow(`
		SELECT name, email FROM users
		WHERE id = 1
	`)

	err = row.Scan(&name, &email)
	if err == sql.ErrNoRows {
		fmt.Println("404")
	}
	if err != nil {
		panic(err)
	}
	fmt.Println(name, email)
}

func (uc UserController) Index(w http.ResponseWriter, r *http.Request) {
	db := getDatabase()
	rows, err := db.Query(`
		SELECT * FROM users LIMIT 10
	`)

	if err != nil {
		panic(err)
	}

	var users []User

	for rows.Next() {
		var user User

		err = rows.Scan(&user.id, &user.name, &user.email, &user.password)

		if err != nil {
			panic(err)
		}

		users = append(users, user)
	}

	fmt.Fprintln(w, users)
}

func getDatabase() *sql.DB {
	config := PostgresConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "root",
		Password: "root",
		Database: "unsplash",
		SSLMode:  false,
	}
	db, err := sql.Open("pgx", config.toString())
	if err != nil {
		panic(err)
	}
	// TODO: Will be close immediately
	//defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected!")

	return db
}
