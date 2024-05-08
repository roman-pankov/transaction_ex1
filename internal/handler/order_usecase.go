package handler

import (
	"time"
	"transaction_ex1/internal/entity"
	"transaction_ex1/internal/repo"
)

type OrderUsecase struct {
	productRepo *repo.ProductRepo
	userRepo    *repo.UserRepo
	orderRepo   *repo.OrderRepo
}

func NewOrderUsecase(
	productRepo *repo.ProductRepo,
	userRepo *repo.UserRepo,
	orderRepo *repo.OrderRepo,
) *OrderUsecase {
	return &OrderUsecase{
		productRepo: productRepo,
		userRepo:    userRepo,
		orderRepo:   orderRepo,
	}
}

func (u *OrderUsecase) MakeOrder(userId int, productId int, processingTime int) error {
	// Ищем пользователя
	user, err := u.userRepo.FindUser(userId)
	if err != nil {
		return err
	}

	// Ищем продукт
	product, err := u.productRepo.FindProduct(productId)
	if err != nil {
		return err
	}

	// Создаём заказ
	newOrderId, err := u.orderRepo.GetNextId()
	if err != nil {
		return err
	}
	order := entity.NewOrder(newOrderId, productId, userId)
	err = u.orderRepo.Save(order)
	if err != nil {
		return err
	}

	// Имитируем тормозжение кода
	time.Sleep(time.Duration(processingTime) * time.Second)

	// Списываем с баланса деньги
	err = user.SubtractAmount(product.GetPrice())
	if err != nil {
		return err
	}

	// Сохраняем пользователя
	err = u.userRepo.Save(user)
	if err != nil {
		return err
	}

	return nil
}
