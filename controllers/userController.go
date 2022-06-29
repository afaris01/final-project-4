package controllers

import (
	"final-project-4/helpers"
	"final-project-4/models/inputs"
	"final-project-4/models/responses"
	"final-project-4/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *userController {
	return &userController{userService}
}

func (h *userController) RegisterUser(c *gin.Context) {
	var input inputs.RegisterUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		error_message := gin.H{
			"error": err.Error(),
		}

		resp := helpers.APIResponse("error", error_message)

		c.JSON(http.StatusBadRequest, resp)
		return
	}

	newUser, err := h.userService.CreateUser(input)

	if err != nil {
		error_message := gin.H{
			"error": err.Error(),
		}

		resp := helpers.APIResponse("error", error_message)

		c.JSON(http.StatusBadRequest, resp)
		return
	}

	// success created user
	userResponse := responses.UserRegisterResponse{
		ID:        newUser.ID,
		FullName:  newUser.FullName,
		Email:     newUser.Email,
		CreatedAt: newUser.CreatedAt,
	}

	resp := helpers.APIResponse("success", userResponse)
	c.JSON(http.StatusCreated, resp)

}

func (h *userController) Login(c *gin.Context) {
	var input inputs.LoginUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errorMessages := gin.H{
			"errors": err.Error(),
		}

		response := helpers.APIResponse("failed", errorMessages)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// send to services
	// get user by email
	user, err := h.userService.GetUserByEmail(input.Email)

	if err != nil {
		errorMessages := gin.H{
			"errors": err.Error(),
		}

		response := helpers.APIResponse("failed", errorMessages)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// return when user not found!
	if user.ID == 0 {
		errorMessages := "User not found!"
		response := helpers.APIResponse("failed", errorMessages)
		c.JSON(http.StatusNotFound, response)
		return
	}

	comparePass := helpers.ComparePass([]byte(user.Password), []byte(input.Password))

	if !comparePass {
		response := helpers.APIResponse("failed", "password not match!")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// lets create token!
	jwtService := helpers.NewService()
	token, err := jwtService.GenerateToken(user.ID)

	// return the token!
	response := helpers.APIResponse("ok", gin.H{
		"token": token,
	})
	c.JSON(http.StatusOK, response)
	return
}

func (h *userController) TopUpSaldo(c *gin.Context) {
	var input inputs.TopUpSaldoInput
	currentUser := c.MustGet("currentUser").(int)

	err := c.ShouldBindJSON(&input)

	if err != nil {
		// errors := helpers.FormatValidationError(err)
		errorMessages := gin.H{
			"errors": err.Error(),
		}

		response := helpers.APIResponse("failed", errorMessages)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userTopup, err := h.userService.TopUp(input.Balance, currentUser)

	if err != nil {
		// errorMessages := helpers.FormatValidationError(err)
		response := helpers.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helpers.APIResponse("ok", gin.H{
		"message": fmt.Sprintf("Your balance has been successfully updated to Rp. %d", userTopup.Balance),
	})
	c.JSON(http.StatusOK, response)
	return

}

func (h *userController) UpdateUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(int)

	var inputUserUpdate inputs.UpdateUserInput

	err := c.ShouldBindJSON(&inputUserUpdate)

	if err != nil {
		// errorMessages := helpers.FormatValidationError(err)
		response := helpers.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	_, err = h.userService.UpdateUser(currentUser, inputUserUpdate)

	if err != nil {
		// errorMessages := helpers.FormatValidationError(err)
		response := helpers.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userUpdated, err := h.userService.GetUserByID(currentUser)

	if err != nil {
		response := helpers.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userResponse := responses.UserUpdateResponse{
		ID:        userUpdated.ID,
		FullName:  userUpdated.FullName,
		Email:     userUpdated.Email,
		UpdatedAt: userUpdated.UpdatedAt,
	}

	response := helpers.APIResponse("ok", userResponse)
	c.JSON(http.StatusOK, response)
	return
}

func (h *userController) DeleteUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(int)

	_, err := h.userService.DeleteUser(currentUser)

	if err != nil {
		// errorMessages := helpers.FormatValidationError(err)
		response := helpers.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helpers.APIResponse("ok", "Your account has been successfully deleted!")
	c.JSON(http.StatusOK, response)
	return
}
