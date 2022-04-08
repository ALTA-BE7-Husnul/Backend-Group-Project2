package order

import (
	_helper "group-project-2/delivery/helper"
	_entities "group-project-2/entities"

	"gorm.io/gorm"
)

type OrderRepository struct {
	database *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		database: db,
	}
}

func (or *OrderRepository) PostOrder(order _helper.OrderRequestFormat, idToken int) (_entities.Transaction, error) {
	var transaction _entities.Transaction
	transaction.Cart_ID = order.Cart_ID
	transaction.User_ID = uint(idToken)
	transaction.Address = order.Address
	transaction.Payment = order.Payment
	transactionTx := or.database.Save(&transaction)

	if transactionTx.Error != nil {
		return _entities.Transaction{}, transactionTx.Error
	}

	return transaction, nil
}

func (or *OrderRepository) GetOrder(idToken int) ([]_entities.Transaction, error) {
	var transaction []_entities.Transaction
	tx := or.database.Where("user_id = ?", idToken).Find(&transaction)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return transaction, nil
}
