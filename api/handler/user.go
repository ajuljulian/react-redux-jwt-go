package handler

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/ajuljulian/react-jwt-go/model"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Register reads registration parameters from the request and persists a new user in the database.
func (h *Handler) Register(c echo.Context) (err error) {
	// Bind the request to a new User object
	u := new(model.User)
	if err = c.Bind(u); err != nil {
		log.Println("could not bind request to a new User object")
		return echo.ErrBadRequest
	}

	// We want to store a hashed version of the password in the database
	if u.Password, err = hashPassword(u.Password); err != nil {
		log.Println(err)
		return echo.ErrInternalServerError
	}

	role := model.Role{Name: "ROLE_USER"}
	u.Roles = append(u.Roles, role)

	h.DB.Create(&u)

	log.Printf("created user with username: %s\n", u.Username)

	r := struct {
		Message string `json:"message"`
	}{
		Message: "User registered successfully!!",
	}

	return c.JSON(http.StatusOK, r)
}

// Login checks the provided credentials and, if they match, returns a JWT token.
func (h *Handler) Login(c echo.Context) (err error) {
	// Bind the request to a new Usr object
	u := new(model.User)
	if err = c.Bind(u); err != nil {
		log.Println("could not bind request to a new User object")
		return echo.ErrBadRequest
	}

	var user model.User
	result := h.DB.Where("username = ?", u.Username).Preload("Roles").First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Println("user not found")
		return echo.ErrUnauthorized
	}

	// Check to see if the hashed password matches the one that was previously persisted.
	passwordsMatch := comparePasswords(user.Password, u.Password)
	if !passwordsMatch {
		log.Println("passwords don't match")
		return echo.ErrUnauthorized
	}

	// Create JWT token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it in the response
	user.Token, err = token.SignedString([]byte(Key))
	if err != nil {
		log.Println("could not generate encoded token")
		return echo.ErrInternalServerError
	}

	var roles []string
	for _, role := range user.Roles {
		roles = append(roles, role.Name)
	}

	r := struct {
		Token    string   `json:"token"`
		Username string   `json:"username"`
		Roles    []string `json:"roles"`
		ID       uint     `json:"id"`
		Email    string   `json:"email"`
	}{
		Token:    user.Token,
		Username: user.Username,
		Roles:    roles,
		ID:       user.ID,
		Email:    user.Email,
	}

	return c.JSON(http.StatusOK, r)
}

// hashPassword hashes the specified password and returns it, or returns an error if one occurs
func hashPassword(pwd string) (string, error) {
	byteHash := []byte(pwd)
	hashedPwd, err := bcrypt.GenerateFromPassword(byteHash, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return "", errors.New("could not generate hash from password")
	}

	return string(hashedPwd), nil
}

// comparePasswords compares a hashed password to a plaintext password.
// It returns true if they match, false otherwise
func comparePasswords(hashedPassword string, plainPassword string) bool {
	byteHashPassword := []byte(hashedPassword)
	bytePlainPassword := []byte(plainPassword)

	err := bcrypt.CompareHashAndPassword(byteHashPassword, bytePlainPassword)
	if err != nil {
		log.Printf("hashed and plain passwords don't match: %v\n", err)
		return false
	}

	return true
}
