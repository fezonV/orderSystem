package domain

import "errors"

var (
	ErrOrderAlreadyPaid     = errors.New("нельзя изменить оплаченный заказ")
	ErrOrderAlreadyCanceled = errors.New("нельзя изменить отмененный заказ")
	ErrOrderIsEmpty         = errors.New("нельзя оплатить пустой заказ")
	ErrInvalidQuantity      = errors.New("количество должно быть положительным")
	ErrInvalidPrice         = errors.New("цена не может быть отрицательной")
	ErrInvalidID            = errors.New("id не может быть отрицательным")
	ErrOrderNotFound        = errors.New("заказ не найден")
)
