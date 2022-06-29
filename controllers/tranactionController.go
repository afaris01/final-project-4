package controllers

import (
	"final-project-4/helpers"
	"final-project-4/models/inputs"
	"final-project-4/models/responses"
	"final-project-4/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionHistoryController struct {
	transactionHistoryService services.TransactionHistoryService
	userService               services.UserService
	productService            services.ProductService
}

func NewTransactionHistoryController(transactionHistoryService services.TransactionHistoryService, userService services.UserService, productService services.ProductService) *transactionHistoryController {
	return &transactionHistoryController{transactionHistoryService, userService, productService}
}

func (h *transactionHistoryController) NewTransaction(c *gin.Context) {
	var input inputs.InputTransaction

	err := c.ShouldBindJSON(&input)

	currentUser := c.MustGet("currentUser").(int)

	if err != nil {
		resp := helpers.APIResponse("error", err.Error())
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	newTransaction, err := h.transactionHistoryService.CreateTransaction(input, currentUser)

	if err != nil {
		resp := helpers.APIResponse("error", err.Error())
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	product, err := h.productService.GetProductByID(input.ProductID)

	billResponse := responses.NewTransactionBillResponse{
		TotalPrice:   newTransaction.TotalPrice,
		Quantity:     input.Quantity,
		ProductTitle: product.Title,
	}

	newTransactionResponse := responses.NewTransactionResponse{
		Message:         "You have successfully purchased the product",
		TransactionBill: billResponse,
	}

	resp := helpers.APIResponse("success", newTransactionResponse)
	c.JSON(http.StatusCreated, resp)
}

func (h *transactionHistoryController) GetMyTransaction(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(int)

	err := c.ShouldBind(&currentUser)

	if err != nil {
		errorMessages := gin.H{
			"errors": err.Error(),
		}

		response := helpers.APIResponse("failed", errorMessages)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	transactions, err := h.transactionHistoryService.GetMyTransaction(currentUser)

	if err != nil {
		errorMessages := gin.H{
			"errors": err.Error(),
		}

		response := helpers.APIResponse("failed", errorMessages)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helpers.APIResponse("success", transactions)
	c.JSON(http.StatusOK, response)
	return
}

func (h *transactionHistoryController) GetUserTransaction(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(int)

	err := c.ShouldBind(&currentUser)

	if err != nil {
		errorMessages := gin.H{
			"errors": err.Error(),
		}

		response := helpers.APIResponse("failed", errorMessages)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	transactions, err := h.transactionHistoryService.GetUserTransaction(currentUser)

	if err != nil {
		errorMessages := gin.H{
			"errors": err.Error(),
		}

		response := helpers.APIResponse("failed", errorMessages)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helpers.APIResponse("success", transactions)
	c.JSON(http.StatusOK, response)
	return

}
