package order

import (
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

func (or *OrderRepository) PostOrder(order _entities.Transaction, orderCartID []uint) (_entities.Transaction, int, error) {
	tx := or.database.Save(&order)
	if tx.Error != nil {
		return order, 0, tx.Error
	}
	for i := range orderCartID {
		var carts _entities.Cart
		var transactionDetail _entities.TransactionDetail
		txFindCart := or.database.Find(&carts, orderCartID[i])
		if txFindCart.Error != nil {
			return order, 0, txFindCart.Error
		}
		transactionDetail.Transaction_ID = order.ID
		transactionDetail.Product_ID = carts.Product_ID
		transactionDetail.Qty = carts.Qty
		transactionDetail.Total = carts.Total

		txTransactionDetail := or.database.Save(&transactionDetail)
		if txTransactionDetail.Error != nil {
			return order, 0, txTransactionDetail.Error
		}
		var products _entities.Product
		txFindProduct := or.database.Find(&products, carts.Product_ID)
		if txFindProduct.Error != nil {
			return order, 0, txFindProduct.Error
		}
		products.Qty -= carts.Qty

		txUpdateProduct := or.database.Save(&products)
		if txUpdateProduct.Error != nil {
			return order, 0, txUpdateProduct.Error
		}

		deleteErr := or.database.Unscoped().Delete(&_entities.Cart{}, orderCartID[i])
		if deleteErr.Error != nil {
			return order, 0, deleteErr.Error
		}
	}

	txAddress := or.database.Save(&order.Address)
	if txAddress.Error != nil {
		return order, 0, txAddress.Error
	}

	txPayment := or.database.Save(&order.Payment)
	if txPayment.Error != nil {
		return order, 0, txPayment.Error
	}
	return order, 1, nil
}

func (or *OrderRepository) GetOrder(idToken int) ([]_entities.Transaction, error) {
	var transaction []_entities.Transaction
	tx := or.database.Where("user_id = ?", idToken).Find(&transaction)
	if tx.Error != nil {
		return nil, tx.Error
	}
	for i := range transaction {
		var address []_entities.Address
		txAddress := or.database.Where("id = ?", transaction[i].ID).Find(&address)
		if txAddress.Error != nil {
			return nil, txAddress.Error
		}
		transaction[i].Address.ID = address[i].ID
		transaction[i].Address.Street = address[i].Street
		transaction[i].Address.City = address[i].City
		transaction[i].Address.State = address[i].State
		transaction[i].Address.Zipcode = address[i].Zipcode

		var payment []_entities.Payment
		txPayment := or.database.Where("id = ?", transaction[i].ID).Find(&payment)
		if txPayment.Error != nil {
			return nil, txPayment.Error
		}
		transaction[i].Payment.ID = payment[i].ID
		transaction[i].Payment.Method = payment[i].Method
		transaction[i].Payment.DestinationBank = payment[i].DestinationBank
	}
	return transaction, nil
}
