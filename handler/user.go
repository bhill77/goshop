package handler

import (
	"errors"
	"net/http"

	"github.com/bhill77/goshop/config"
	"github.com/bhill77/goshop/entity"
	"github.com/bhill77/goshop/middleware"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserHandler struct {
	db   *gorm.DB
	conf config.Config
}

func NewUserHandler(db *gorm.DB, conf config.Config) UserHandler {
	return UserHandler{
		db:   db,
		conf: conf,
	}
}

func (h UserHandler) Index(c echo.Context) error {
	var users []entity.User
	h.db.Find(&users)
	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "data": users})
}

func (h UserHandler) Create(c echo.Context) error {
	var user entity.User

	e := validateRequest(c, &user)
	if len(e) > 0 {
		err := map[string]interface{}{"validationError": e}
		return c.JSON(http.StatusBadRequest, err)
	}

	user.Password, _ = HashPassword(user.Password)
	err := h.db.Create(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, "gagal insert")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "data": user})
}

func (h UserHandler) Update(c echo.Context) error {
	id := c.Param("id")

	var user entity.User
	err := h.db.Where("id", id).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, "Invalid id")
	}

	var payload entity.User

	e := validateRequest(c, &payload)
	if len(e) > 0 {
		err := map[string]interface{}{"validationError": e}
		return c.JSON(http.StatusBadRequest, err)
	}

	user.Name = payload.Name
	user.Address = payload.Address
	user.PhoneNumber = payload.PhoneNumber

	if err := h.db.Updates(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "gagal update")
	}

	return c.JSON(http.StatusOK, user)
}

func (h UserHandler) Delete(c echo.Context) error {
	id := c.Param("id")

	var user entity.User
	err := h.db.Where("id", id).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, "Invalid id")
	}

	if err := h.db.Delete(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "gagal hapus")
	}

	return c.JSON(http.StatusOK, "user berhasil dihapus")
}

func (h UserHandler) Show(c echo.Context) error {
	id := c.Param("id")

	var user entity.User
	err := h.db.Where("id", id).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, "Invalid id")
	}

	return c.JSON(http.StatusOK, user)
}

func (h UserHandler) Login(c echo.Context) error {
	var payload entity.LoginRequest
	c.Bind(&payload)

	var user entity.User
	err := h.db.Where("email", payload.Email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusBadRequest, "Akun tidak ditemukan")
	}

	isValid := CheckPasswordHash(payload.Password, user.Password)
	if !isValid {
		return c.JSON(http.StatusBadRequest, "Email atau password salah")
	}

	// create jwt token
	claims := &middleware.JwtCustomClaims{
		ID: user.ID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(h.conf.JwtSecret))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})

}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
