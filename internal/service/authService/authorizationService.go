package authService

import (
	"GeoServiseAppDate/internal/models"
	"GeoServiseAppDate/internal/repository/authRepository"
	"fmt"
	"github.com/go-chi/jwtauth"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	SaveUser(user models.User) error
	GetToken(user models.User) (string, error)
}

type authService struct {
	r authRepository.AuthRepository
}

func NewAuthService(r authRepository.AuthRepository) authService {
	return authService{r}
}

func (as *authService) SaveUser(user models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	if err := as.r.SaveUser(models.User{
		Login:    user.Login,
		Password: string(hashedPassword),
	}); err != nil {
		return err
	}
	return nil
}

func (as *authService) GetToken(user models.User) (string, error) {
	var tokenAuth = jwtauth.New("HS256", []byte("mySecret"), nil)
	userInBD, err := as.r.GetUser(user)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userInBD.Password), []byte(user.Password)); err != nil {
		return "", err
	}

	_, tokenString, err := tokenAuth.Encode(map[string]interface{}{"username": userInBD.Login})
	return tokenString, err
}

type authServiceProxy struct {
	realObject authService
}

func NewAuthServiceProxy(r authRepository.AuthRepository) AuthService {
	return &authServiceProxy{NewAuthService(r)}
}

func (asp *authServiceProxy) SaveUser(user models.User) error {
	isExist, err := asp.realObject.r.CheckUser(user)
	if err != nil {
		return err
	} else if isExist {
		return fmt.Errorf("the user is already registered")
	}

	return asp.realObject.SaveUser(user)
}

func (asp *authServiceProxy) GetToken(user models.User) (string, error) {
	if isExist, err := asp.realObject.r.CheckUser(user); err != nil {
		return "", err
	} else if !isExist {
		return "", fmt.Errorf("the user is not exist")
	}
	return asp.realObject.GetToken(user)
}
