package models

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Service is our main struct, including the database instance for working with data.
type Service struct {

	// db is an instance of the SQLite database.
	db *gorm.DB
}

// NewService is the constructor for the Service struct.
func NewService(db *gorm.DB) (*Service, error) {

	// We check if the database instance is nil, which would indicate an issue.
	if db == nil {
		return nil, errors.New("please provide a valid connection")
	}

	// We initialize our service with the passed database instance.
	s := &Service{db: db}
	return s, nil
}

// CreateUser is a method that creates a new user record in the database.
func (s *Service) CreateUser(ctx context.Context, nu NewUser, now time.Time) (User, error) {

	// We hash the user's password for storage in the database.
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(nu.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, fmt.Errorf("generating password hash: %w", err)
	}

	// We prepare the User record.
	u := User{
		Name:         nu.Name,
		Email:        nu.Email,
		PasswordHash: string(hashedPass),
		DateCreated:  now,
		DateUpdated:  now,
	}

	// We attempt to create the new User record in the database.
	err = s.db.Create(&u).Error
	if err != nil {
		return User{}, err
	}

	// Successfully created the record, return the user.
	return u, nil
}

// Authenticate is a method that checks a user's provided email and password against the database.
func (s *Service) Authenticate(ctx context.Context, email, password string, now time.Time) (jwt.RegisteredClaims,
	error) {

	// We attempt to find the User record where the email
	// matches the provided email.
	var u User
	tx := s.db.Where("email = ?", email).First(&u)
	if tx.Error != nil {
		return jwt.RegisteredClaims{}, tx.Error
	}

	// We check if the provided password matches the hashed password in the database.
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	if err != nil {
		return jwt.RegisteredClaims{}, err
	}

	// Successful authentication! Generate JWT claims.
	c := jwt.RegisteredClaims{
		Issuer:    "service project",
		Subject:   strconv.FormatUint(uint64(u.ID), 10),
		Audience:  jwt.ClaimStrings{"students"},
		ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour)),
		IssuedAt:  jwt.NewNumericDate(now),
	}

	// And return those claims.
	return c, nil
}
