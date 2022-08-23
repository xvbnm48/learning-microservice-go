package data

import (
	"context"
	"database/sql"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const dbTimeOut = time.Second * 3

var db *sql.DB

// New: for used to create instance of data package
// Model, which embeds all the types we want to be available to our application.
func New(dbPoll *sql.DB) Models {
	db = dbPoll

	return Models{
		User: User{},
	}
}

// Models is the type for this package. Note that any model that is included as a member
// in this type is available to us throughout the application, anywhere that the
// app variable is used, provided that the model is also added in the New function.
type Models struct {
	User User
}

type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	FirtsName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Password  string    `json:"password"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// GetAll returns all the users in the database, sorted by last name
func (u *User) GetAll() ([]*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := `SELECT id, email, firts_name, last_name, password, active, created_at, updated_at FROM users ORDER BY last_name`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []*User

	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.FirtsName,
			&user.LastName,
			&user.Password,
			&user.Active,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

		if err != nil {
			log.Println("error scanning", err)
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}

// GetByEmail: returns a one user by email
func (u *User) GetByEmail(email string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := `SELECT id, email, firts_name, last_name, password, active, created_at, updated_at FROM users WHERE email = $1`

	var user User
	row := db.QueryRowContext(ctx, query, email)

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FirtsName,
		&user.LastName,
		&user.Password,
		&user.Active,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil

}

// GetOne: return one user by id
func (u *User) GetOne(id string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := `SELECT id, email, firts_name, last_name, password, active, created_at, updated_at FROM users WHERE id = $1`

	var user User
	row := db.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FirtsName,
		&user.LastName,
		&user.Password,
		&user.Active,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Update: updates one user in the database, using the information
// stored in the receiver u
func (u *User) Update() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	stmt := `UPDATE users set email = $1, firts_name = $2, last_name = $3, user_active = $4, updated_at = $5 WHERE id = $6`

	_, err := db.ExecContext(ctx, stmt, u.Email, u.FirtsName, u.LastName, u.Active, time.Now(), u.ID)
	if err != nil {
		return err
	}

	return nil
}

// Delete: deletes one user from the database, by user id
func (u *User) Delete() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	stmt := `DELETE FROM users WHERE id = $1`
	_, err := db.ExecContext(ctx, stmt, u.ID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteById deletes one user from the database, by id
func (u *User) DeleteById(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	stmt := `DELETE FROM users WHERE id = $1`
	_, err := db.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	return nil
}

// Insert: Inserts a new user into the database, and returns the ID of the newly inserted row
func (u *User) Insert(user User) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return 0, err
	}

	var newId int
	stmt := `INSERT INTO users (email, firts_name, last_name, password, active, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	err = db.QueryRowContext(ctx, stmt, user.Email, user.FirtsName, user.LastName, hashedPassword, user.Active, time.Now(), time.Now()).Scan(&newId)

	if err != nil {
		return 0, err
	}
	return newId, nil
}
