package controllers

import (
	"net/http"
	"final-project-4/helpers"
	"final-project-4/models/inputs"
	"final-project-4/models/responses"
	"final-project-4/services"

	"github.com/gin-gonic/gin"
)

type categoryController struct {
	categoryService services.CategoryService
	userService     services.UserService
}

func NewCategoryController(categoryService services.CategoryService, userService services.UserService) *categoryController {
	return &categoryController{categoryService, userService}
}

func (h *categoryController) CreateCategory(c *gin.Context) {
	var input inputs.CreateCategory

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

	newCategory, err := h.categoryService.CreateCategory(input)

	if err != nil {
		error_message := gin.H{
			"error": err.Error(),
		}

		resp := helpers.APIResponse("error", error_message)

		c.JSON(http.StatusBadRequest, resp)
		return
	}

	// success created user
	Response := responses.CreateCategoryResponse{
		ID:        newCategory.ID,
		Type:      newCategory.Type,
		CreatedAt: newCategory.CreatedAt,
	}

	resp := helpers.APIResponse("success", Response)
	c.JSON(http.StatusCreated, resp)

}

func (h *categoryController) UpdateCategory(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(int)

	var inputUpdate inputs.UpdateCategory
	var IDInput inputs.IDCategory

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

	_, err = h.categoryService.UpdateCategory(IDInput.ID, inputUpdate)

	if err != nil {
		// errorMessages := helpers.FormatValidationError(err)
		response := helpers.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	categoryUpdated, err := h.categoryService.GetCategoryByID(currentUser)

	if err != nil {
		response := helpers.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userResponse := responses.UpdateCategoryResponse{
		ID:        categoryUpdated.ID,
		Type:      categoryUpdated.Type,
		UpdatedAt: categoryUpdated.UpdatedAt,
	}

	response := helpers.APIResponse("ok", userResponse)
	c.JSON(http.StatusOK, response)
	return
}

func (h *categoryController) DeleteCategory(c *gin.Context) {

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

	_, err = h.categoryService.DeleteCategory(Deleted.ID)

	if err != nil {
		response := helpers.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	message := "Category deleted"

	response := helpers.APIResponse("ok", message)
	c.JSON(http.StatusOK, response)
}

func (h *categoryController) GetAllCategory(c *gin.Context) {

	allCategory, err := h.categoryService.GetCategory()

	if err != nil {
		error_message := gin.H{
			"error": err.Error(),
		}

		resp := helpers.APIResponse("error", error_message)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp := helpers.APIResponse("success", allCategory)
	c.JSON(http.StatusOK, resp)
}
