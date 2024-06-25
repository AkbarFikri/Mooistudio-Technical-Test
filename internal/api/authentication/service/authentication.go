package service

import (
	"database/sql"
	"errors"
	"github.com/AkbarFikri/Mooistudio-Technical-Test/internal/api/authentication/dto"
	"github.com/AkbarFikri/Mooistudio-Technical-Test/internal/api/authentication/repository"
	"github.com/AkbarFikri/Mooistudio-Technical-Test/internal/domain"
	customErr "github.com/AkbarFikri/Mooistudio-Technical-Test/internal/pkg/error"
	"golang.org/x/net/context"
	"log"
	"time"
)

type AuthService interface {
	Register(ctx context.Context, req dto.AuthRequest) (dto.AuthResponse, error)
	Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error)
}

type authService struct {
	AuthRepository repository.AuthRepository
}

func NewAuthService(ar repository.AuthRepository) AuthService {
	return &authService{ar}
}

func (s authService) Register(ctx context.Context, req dto.AuthRequest) (dto.AuthResponse, error) {
	count, err := s.AuthRepository.CountEmail(ctx, req.Email)
	if err != nil {
		log.Printf("error repository not recognized %v", err)
		return dto.AuthResponse{}, customErr.ErrorGeneral
	}

	if count > 0 {
		log.Printf("error email already exists")
		return dto.AuthResponse{}, customErr.ErrorEmailAlreadyUsed
	}

	user := domain.User{
		Email:    req.Email,
		Password: req.Password,
		Fullname: req.FullName,
	}
	user.Create()

	if err := s.AuthRepository.Save(ctx, user); err != nil {
		log.Printf("error repository when create user %v", err)
		return dto.AuthResponse{}, customErr.ErrorGeneral
	}

	return dto.AuthResponse{
		FullName: user.Fullname,
		Email:    user.Email,
		ID:       user.ID,
	}, nil
}

func (s authService) Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error) {
	user, err := s.AuthRepository.FindUserByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("invalid user email")
			return dto.LoginResponse{}, customErr.ErrorInvalidEmailOrPassword
		} else {
			log.Printf("error repository not recognized %v", err)
			return dto.LoginResponse{}, customErr.ErrorGeneral
		}
	}

	if !user.ComparePassword(req.Password) {
		log.Printf("error password not match")
		return dto.LoginResponse{}, customErr.ErrorInvalidEmailOrPassword
	}

	token := user.CreateAccessToken()

	return dto.LoginResponse{
		Token:     token,
		ExpiredAt: time.Now().Add(24 * time.Hour).UnixNano(),
	}, nil
}
