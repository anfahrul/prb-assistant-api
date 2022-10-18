package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/anfahrul/prb-assistant-api/entity"
	"github.com/anfahrul/prb-assistant-api/utils/token"
	"golang.org/x/crypto/bcrypt"
)

type userRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepositoryImpl{DB: db}
}

func (repository *userRepositoryImpl) Insert(ctx context.Context, user entity.User) (entity.User, error) {
	script := "INSERT INTO user(username, password, email, role, isLoggedIn) VALUES (?, ?, ?, ?, ?)"

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}

	_, err = repository.DB.ExecContext(ctx, script, user.Username, hashedPassword, user.Email, user.Role, user.IsLoggedIn)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repository *userRepositoryImpl) LoginCheck(ctx context.Context, user entity.User) (string, error) {
	script := `SELECT id, username, password, email, role, isLoggedIn FROM user WHERE username=?`
	rows, err := repository.DB.QueryContext(ctx, script, user.Username)
	u := entity.User{}
	var role string
	if err != nil {
		return "", err
	}
	defer rows.Close()

	if rows.Next() {
		// ada
		rows.Scan(
			&u.Id,
			&u.Username,
			&u.Password,
			&u.Email,
			&role,
			&u.IsLoggedIn,
		)

		err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password))
		if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
			return "", err
		}

		token, err := token.GenerateToken(u.Username, role)
		if err != nil {
			fmt.Println("error", err.Error())
			return "", err
		}

		return token, nil
	}

	return "", err
}

func (repository *userRepositoryImpl) UpdateLoginStatus(ctx context.Context, user entity.User, loginStatus int) error {
	script := `
		UPDATE user
		SET isLoggedIn = ?
		WHERE username = ?
	`

	_, err := repository.DB.ExecContext(ctx, script, loginStatus, user.Username)
	if err != nil {
		return err
	}

	return nil
}
