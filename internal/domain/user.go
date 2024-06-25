package domain

import (
	"github.com/AkbarFikri/Mooistudio-Technical-Test/internal/pkg/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type User struct {
	ID        string    `db:"id"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Fullname  string    `db:"full_name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (u *User) HashPassword(password string) (string, error) {
	pass := []byte(password)
	result, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

func (u *User) Create() {
	u.ID = uuid.NewString()
	hash, err := u.HashPassword(u.Password)
	if err != nil {
		log.Printf("error when hashing password")
	}
	u.Password = string(hash)
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}

func (u *User) CreateAccessToken() string {
	accessData := map[string]interface{}{
		"id":    u.ID,
		"email": u.Email,
	}
	accessToken, _ := jwt.SignJWT(accessData, 24)
	return accessToken
}
