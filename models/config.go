package models

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	"golang.org/x/crypto/pbkdf2"
)

// LoginType represents the type of login, either anonymous or protected.
type LoginType string

const (
	// AnonymousLoginType is for anonymous login.
	AnonymousLoginType = LoginType("anonymous")

	// ProtectedLoginType is for protected login.
	ProtectedLoginType = LoginType("protected")
)

type Config struct {
	ID        uint      `json:"-" gorm:"primarykey"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

	// Strava APP Credentials
	ClientId     int    `json:"client_id"`
	ClientSecret string `json:"client_secret"`

	// App Login Details
	LoginType     LoginType `json:"-"`
	LoginUsername string    `json:"-"`
	LoginPassword string    `json:"-"`
}

// SetupCompleted checks if the configuration setup is completed
func (c *Config) SetupCompleted() bool {
	return c.ID != 0
}

// hashPassword generates a hashed password using PBKDF2
// with the provided password and secret key.
func (c *Config) hashPassword(password, secretKey string) string {
	hash := pbkdf2.Key([]byte(password), []byte(secretKey), 10000, 32, sha256.New)
	return hex.EncodeToString(hash)
}

// SetPassword hashes the provided password with
// the given secret key and sets it to the Config's LoginPassword field
func (c *Config) SetPassword(password, secretKey string) {
	c.LoginPassword = c.hashPassword(password, secretKey)
}

// CheckLogin verifies the provided username and password against the stored credentials.
func (c *Config) CheckLogin(username, password, secretKey string) bool {
	passwd := c.hashPassword(password, secretKey)
	return c.LoginUsername == username && c.LoginPassword == passwd
}
