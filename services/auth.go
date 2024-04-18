package services

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mexirica/todo/infra"
	"github.com/mexirica/todo/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

var jwtKey = []byte("superSecretKey")

func GenerateAccessToken(username *string, id *uint) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["iss"] = &username
	claims["id"] = &id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	claims["iat"] = time.Now().Unix()
	claims["sub"] = "API Go"

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateRefreshToken(email *string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["iss"] = &email
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()
	claims["iat"] = time.Now().Unix()
	claims["sub"] = "API Go"

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	erro := UpdateRefreshToken(email, &tokenString)
	if erro != nil {
		return "", erro
	}

	return tokenString, nil
}

func UpdateRefreshToken(email, refreshToken *string) error {
	result := infra.DB.Model(&models.User{}).Where("email = ?", email).Update("refresh_token", refreshToken)
	if result.Error != nil {
		return fmt.Errorf("error to update the column refresh token: %v", result.Error)
	}

	return nil
}

func SignUp(c *gin.Context) {
	var dto models.NewUser
	if err := c.BindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read the request body"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate password hash"})
		return
	}
	layout := "02/01/2006"
	birthday, _ := time.Parse(layout, dto.Birthday)
	user := models.User{
		Name:      dto.Name,
		Surname:   dto.Surname,
		Email:     dto.Email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
		Birthday:  birthday,
		Company:   dto.Company,
	}

	if err := infra.DB.Model(&models.User{}).Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func Login(c *gin.Context) {
	var login models.Login
	if err := c.BindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read the request body"})
		return
	}

	var user models.User
	if err := infra.DB.Model(&models.User{}).Where("email = ?", login.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	refreshToken, err := GenerateRefreshToken(&login.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate refresh token"})
		return
	}

	accessToken, err := GenerateAccessToken(&login.Email, &user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate access token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
