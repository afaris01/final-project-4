package services

import (
	"errors"
	"final-project-4/models"
	"final-project-4/models/inputs"
	"final-project-4/models/responses"
	"final-project-4/repositories"
)

type TransactionHistoryService interface {
	CreateTransaction(transactionInput inputs.InputTransaction, IDUser int) (models.TransactionHistory, error)
	GetMyTransaction(IDUser int) ([]responses.TransactionHistoryResponse, error)
	GetUserTransaction(IDUser int) ([]responses.UserTransactionHistoryResponse, error)
}

type transactionHistoryService struct {
	transactionHistoryRepository repositories.TransactionHistoryRepository
	productRepository            repositories.ProductRepository
	userRepository               repositories.UserRepository
}

func NewTransactionHistoryService(transactionHistoryRepository repositories.TransactionHistoryRepository, productRepository repositories.ProductRepository, userRepository repositories.UserRepository) *transactionHistoryService {
	return &transactionHistoryService{transactionHistoryRepository, productRepository, userRepository}
}

func (s *transactionHistoryService) CreateTransaction(transactionInput inputs.InputTransaction, IDUser int) (models.TransactionHistory, error) {
	newTransactionHistory := models.TransactionHistory{}

	newTransactionHistory.ProductID = transactionInput.ProductID
	newTransactionHistory.Quantity = transactionInput.Quantity

	// query product
	product, err := s.productRepository.GetProductByID(transactionInput.ProductID)

	if err != nil {
		return models.TransactionHistory{}, err
	}

	if product.ID == 0 {
		return models.TransactionHistory{}, err
	}

	// ketika jumlah tidak mencukupi
	if product.Stock < transactionInput.Quantity {
		return models.TransactionHistory{}, errors.New("Jumlah tidak mencukupi")
	}

	// query user
	datauser, err := s.userRepository.GetByID(IDUser)

	if err != nil {
		return models.TransactionHistory{}, err
	}

	if datauser.ID == 0 {
		return models.TransactionHistory{}, err
	}

	// saldo tidak mencukupi
	if datauser.Balance < (product.Price * transactionInput.Quantity) {
		return models.TransactionHistory{}, errors.New("Saldo tidak mencukupi")
	}

	// pastikan balance tersedia
	buyAmount := product.Price * transactionInput.Quantity
	newTransactionHistory.TotalPrice = buyAmount
	newTransactionHistory.UserID = IDUser

	// kurangi stock
	productUpdate := models.Product{
		ID:    transactionInput.ProductID,
		Stock: product.Stock - transactionInput.Quantity,
	}

	_, err = s.productRepository.Update(productUpdate)

	if err != nil {
		return models.TransactionHistory{}, err
	}

	// store data ke transactions history
	transactionCreated, err := s.transactionHistoryRepository.Save(newTransactionHistory)

	if err != nil {
		return models.TransactionHistory{}, err
	}

	_, err = s.userRepository.UpdateSaldo(IDUser, buyAmount)

	if transactionCreated.ID == 0 {
		return models.TransactionHistory{}, err
	}

	return transactionCreated, nil
}

func (s *transactionHistoryService) GetMyTransaction(IDUser int) ([]responses.TransactionHistoryResponse, error) {
	myTransaction, err := s.transactionHistoryRepository.GetTransactionByIDUser(IDUser)

	if err != nil {
		return []responses.TransactionHistoryResponse{}, err
	}

	if len(myTransaction) < 1 {
		return []responses.TransactionHistoryResponse{}, err
	}

	var myTransactionResponse []responses.TransactionHistoryResponse

	for _, item := range myTransaction {
		productTemp, _ := s.productRepository.GetProductByID(item.ProductID)

		temp := responses.TransactionHistoryResponse{
			ID:         item.ID,
			ProductID:  item.ProductID,
			UserID:     item.UserID,
			Quantity:   item.Quantity,
			TotalPrice: item.TotalPrice,
			CreatedAt:  item.CreatedAt,
			UpdatedAt:  item.UpdatedAt,
			Product:    productTemp,
		}
		myTransactionResponse = append(myTransactionResponse, temp)
	}

	return myTransactionResponse, nil
}

func (s *transactionHistoryService) GetUserTransaction(IDUser int) ([]responses.UserTransactionHistoryResponse, error) {

	userdata, err := s.userRepository.GetByID(IDUser)
	if userdata.ID == 0 {
		return []responses.UserTransactionHistoryResponse{}, errors.New("users not found!")
	}

	if userdata.Role != "admin" {
		return []responses.UserTransactionHistoryResponse{}, errors.New("Unauthorized user!")
	}

	myTransaction, err := s.transactionHistoryRepository.GetAllTransaction()

	if err != nil {
		return []responses.UserTransactionHistoryResponse{}, err
	}

	if len(myTransaction) < 1 {
		return []responses.UserTransactionHistoryResponse{}, err
	}

	var myTransactionResponse []responses.UserTransactionHistoryResponse

	for _, item := range myTransaction {
		productTemp, _ := s.productRepository.GetProductByID(item.ProductID)
		userTemp, _ := s.userRepository.GetByID(item.UserID)

		temp := responses.UserTransactionHistoryResponse{
			ID:         item.ID,
			ProductID:  item.ProductID,
			UserID:     item.UserID,
			Quantity:   item.Quantity,
			TotalPrice: item.TotalPrice,
			CreatedAt:  item.CreatedAt,
			UpdatedAt:  item.UpdatedAt,
			Product:    productTemp,
			User:       userTemp,
		}
		myTransactionResponse = append(myTransactionResponse, temp)
	}

	return myTransactionResponse, nil
}
