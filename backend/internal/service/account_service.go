package service

import (
	"github.com/gin-gonic/gin"
	"github.com/kmj36/fieldnotes-tech-blog/internal/dto"
	"github.com/kmj36/fieldnotes-tech-blog/internal/model"
	"github.com/kmj36/fieldnotes-tech-blog/internal/repository"
	"github.com/kmj36/fieldnotes-tech-blog/pkg/cryption"
)

// 관리자 계정 관련 비즈니스 로직
type AccountService struct {repo *repository.AccountRepository}

func NewAccountService(repo *repository.AccountRepository) *AccountService {
	return &AccountService{repo: repo}
}

func (s *AccountService) Create(ctx *gin.Context, req *dto.CreateAccountRequest) error {

	if existing, _ := s.repo.FindByAccountID(req.AccountID) ; existing != nil {
		return dto.ErrAccountAlreadyExists
	} 

	var hash string
	var err error

	if hash, err = cryption.HashPassword(req.Password) ; err != nil {
		return err
	}

	newAccount := &model.Account{
		AccountID: req.AccountID,
		PasswordHash: hash,
		Nickname: req.Nickname,
		AvatarURL: req.AvatarURL,
		Role: "ADMIN",
		Status: "ACTIVE",
	}

	return s.repo.Create(newAccount)
}