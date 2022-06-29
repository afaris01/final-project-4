package controllers

import (
	"net/http"
	"final-project-4/helpers"
	"final-project-4/models/inputs"
	"final-project-4/models/responses"
	"final-project-4/services"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productService services.ProductService
	userService    services.UserService
}

func NewProductController(productService services.ProductService, userService services.UserService) *productController {
	return &productController{productService, userService}
}

func (h *productController) CreateProduct(c *gin.Context) {
	var input inputs.CreateProduct

	err := c.ShouldBindJSON(&input)

	if err != nil {

		resp := helpers.APIResponse("error", err.Error())

		c.JSON(http.StatusBadRequest, resp)
		return
	}

	currentUser := c.MustGet("currentUser").(int)

	userResult, err := h.userService.GetUserByID(currentUser)

	if userResult.Role != "admin" {
		resp := helpers.APIResponse("error", "Unauthorized User!")
		c.JSON(http.StatusUnauthorized, resp)
		return
	}

	newProduct, err := h.productService.CreateProduct(input)

	if err != nil {
		error_message := gin.H{
			"error": err.Error(),
		}

		resp := helpers.APIResponse("error", error_message)

		c.JSON(http.StatusBadRequest, resp)
		return
	}

	// success created user
	Response := responses.CreateProductResponse{
		ID:        newProduct.ID,
		Title:     newProduct.Title,
		Price:     newProduct.Price,
		Stock:     newProduct.Stock,
		CreatedAt: newProduct.CreatedAt,
	}

	resp := helpers.APIResponse("success", Response)
	c.JSON(http.StatusCreated, resp)

}

func (h *productController) UpdateProduct(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(int)

	var inputUpdate inputs.UpdateProduct
	var IDInput inputs.IDProduct

	err := c.ShouldBindJSON(&inputUpdate)
	err = c.ShouldBindUri(&IDInput)

	if err != nil {
		response := helpers.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userResult, err := h.userService.GetUserByID(currentUser)

	if userResult.Role != "admin" {
		resp := helpers.APIResponse("error", "Unauthorized User!")
		c.JSON(http.StatusUnauthorized, resp)
		return
	}

	_, err = h.productService.UpdateProduct(IDInput.ID, inputUpdate)

	if err != nil {
		// errorMessages := helpers.FormatValidationError(err)
		response := helpers.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	productUpdated, err := h.productService.GetProductByID(IDInput.ID)

	if err != nil {
		response := helpers.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	productResponse := responses.UpdateProductResponse{
		ID:        productUpdated.ID,
		Title:     productUpdated.Title,
		Price:     productUpdated.Price,
		Stock:     productUpdated.Stock,
		UpdatedAt: productUpdated.UpdatedAt,
	}

	response := helpers.APIResponse("ok", productResponse)
	c.JSON(http.StatusOK, response)
	return
}

func (h *productController) DeleteProduct(c *gin.Context) {

	currentUser := c.MustGet("currentUser").(int)

	if currentUser == 0 {
		response := helpers.APIResponse("failed", "unauthorized user")
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	var Deleted inputs.DeleteCategory
	err := c.ShouldBindUri(&Deleted)

	if err != nil {
		response := helpers.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userResult, err := h.userService.GetUserByID(currentUser)

	if userResult.Role != "admin" {
		resp := helpers.APIResponse("error", "Unauthorized User!")
		c.JSON(http.StatusUnauthorized, resp)
		return
	}

	_, err = h.productService.DeleteProduct(Deleted.ID)

	if err != nil {
		response := helpers.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	message := "products deleted"

	response := helpers.APIResponse("ok", message)
	c.JSON(http.StatusOK, response)
}

func (h *productController) GetAllProduct(c *gin.Context) {

	allProduct, err := h.productService.GetProduct()

	if err != nil {
		error_message := gin.H{
			"error": err.Error(),
		}

		resp := helpers.APIResponse("error", error_message)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp := helpers.APIResponse("success", allProduct)
	c.JSON(http.StatusOK, resp)
}

func (h *productController) GetProductByID(c *gin.Context) {

	var idproduct inputs.InputProduct

	product, err := h.productService.GetProductByID(idproduct.ID)

	if err != nil {
		response := helpers.APIResponse("failed", "id must be exist!")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	response := helpers.APIResponse("ok", product)
	c.JSON(http.StatusOK, response)
	return
}
