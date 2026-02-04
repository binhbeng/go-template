package repository

import (
	"context"
	"fmt"

	"github.com/binhbeng/goex/internal/dto"
	"github.com/binhbeng/goex/internal/model/entity"
	"github.com/binhbeng/goex/internal/utils"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (m *UserRepository) TableName() string {
	return "users"
}

func (m *UserRepository) WithTx(tx *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: tx,
	}
}

func (m *UserRepository) Update(ctx context.Context, id int, user entity.User) (entity.User, error) {
	if err := m.DB.WithContext(ctx).Save(&user).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (m *UserRepository) GetUserById(ctx context.Context, id int) (entity.User, error) {
	var user entity.User
	if err := m.DB.WithContext(ctx).First(&user, id).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (m *UserRepository) GetListUser(ctx context.Context, filter dto.QueryUsersInput) ([]entity.User, error) {
	var users []entity.User

	query := m.DB.WithContext(ctx).Model(&entity.User{})

	err := query.Scopes(utils.Paginate(filter.PageOptionsDto)).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (m *UserRepository) AddBalance(ctx context.Context, userID int, amount string) (string, error) {
	var user entity.User
	if err := m.DB.WithContext(ctx).
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("id = ?", userID).
		First(&user).Error; err != nil {
		return "", err
	}

	amountDec, err := decimal.NewFromString(amount)
	if err != nil {
		return "", fmt.Errorf("invalid amount format: %w", err)
	}

	currentBalanceDec, err := decimal.NewFromString(user.Balance)
	if err != nil {
		return "", fmt.Errorf("invalid current balance format: %w", err)
	}

	newBalanceDec := currentBalanceDec.Add(amountDec)
	newBalanceStr := newBalanceDec.String()

	err = m.DB.WithContext(ctx).Model(&user).
		Where("id = ?", userID).
		Update("balance", newBalanceStr).Error
	if err != nil {
		return "", err
	}

	return newBalanceStr, nil
}
