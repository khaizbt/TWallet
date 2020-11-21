package handler

import (
	"TWallet/auth"
	"TWallet/helper"
	"TWallet/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct { //Handler manggil service
	userService user.Service
	authService auth.Service
}

func NewuserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {

		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}

		responsError := helper.APIResponse("Create Account Failed", http.StatusUnprocessableEntity, "fail", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, responsError)
		return
	}

	NewUser, err := h.userService.RegisterUser(input)

	if err != nil {
		responsError := helper.APIResponse("Create Account Failed", http.StatusBadRequest, "fail", nil)
		c.JSON(http.StatusBadRequest, responsError)
		return
	}

	token, err := h.authService.GenerateToken(NewUser.ID)
	if err != nil {
		responsError := helper.APIResponse("Create Account Failed", http.StatusBadGateway, "fail", "Unable to send Email")
		c.JSON(http.StatusBadGateway, responsError)
		return
	}

	// Custom Format JSON Response
	formater := user.FormatUser(NewUser, token)
	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formater)

	c.JSON(http.StatusOK, response)
}

//LOGIN
/*

Flow
	User Memasukan Input
	Input Ditangkap Handler
	Mapping Dari Input user ke Input Struct
	Input dari strcut parsing service
	Diservice mencari dengan bantuan repository user dengan email x
	Mencocokan Password
*/

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}

		responsError := helper.APIResponse("Login Failed #LOG001", http.StatusUnprocessableEntity, "fail", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, responsError)
		return
	}

	loggedinUser, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		responsError := helper.APIResponse("Login Failed #LOG002", http.StatusUnprocessableEntity, "fail", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, responsError)
		return
	}

	token, err := h.authService.GenerateToken(loggedinUser.ID)
	if err != nil {
		responsError := helper.APIResponse("Login Failed", http.StatusBadGateway, "fail", "Unable to generate token")
		c.JSON(http.StatusBadGateway, responsError)
		return
	}

	formater := user.FormatUser(loggedinUser, token)

	response := helper.APIResponse("Login Success", http.StatusOK, "success", formater)

	c.JSON(http.StatusOK, response)
}

//Check Email Available
func (h *userHandler) CheckEmailAvailablity(c *gin.Context) {
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}

		responsError := helper.APIResponse("Email Checking Failed", http.StatusUnprocessableEntity, "fail", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, responsError)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": "Something Went Wrong #CHE001"}

		responsError := helper.APIResponse("Email Checking Failed", http.StatusUnprocessableEntity, "fail", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, responsError)
		return
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}

	var metaMessage string

	if isEmailAvailable {
		metaMessage = "Email is Available"
	} else {
		metaMessage = "Email has been registered"
	}

	responsError := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, responsError)

}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar_image") //Keynya

	if err != nil {
		errorMessage := gin.H{"is_uploaded": false}

		responsError := helper.APIResponse("Upload Image Failed #UPA001", http.StatusBadGateway, "fail", errorMessage)
		c.JSON(http.StatusBadGateway, responsError)
		return
	}

	path := "images/avatar/" + file.Filename

	err = c.SaveUploadedFile(file, path)

	if err != nil {
		errorMessage := gin.H{"is_uploaded": false}

		responsError := helper.APIResponse("Upload Image Failed #UPA002", http.StatusBadGateway, "fail", errorMessage)
		c.JSON(http.StatusBadGateway, responsError)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User) //nilai default dari current User adalah interface, jadi harus diubah ke user.user dahulu agar bisa dipparsing
	userID := currentUser.ID
	_, err = h.userService.UploadAvatar(userID, path)

	if err != nil {
		errorMessage := gin.H{"is_uploaded": false}

		responsError := helper.APIResponse("Upload Image Failed #UPA003", http.StatusBadGateway, "fail", errorMessage)
		c.JSON(http.StatusBadGateway, responsError)
		return
	}

	data := gin.H{
		"is_uploaded": true,
	}
	response := helper.APIResponse("Avatar Succesfully Uploaded", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)

}
