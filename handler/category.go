package handler

import (
	"TWallet/category"
	"TWallet/helper"
	"TWallet/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type categoryHandler struct {
	service category.Service
}

func NewCategoryHandler(service category.Service) *categoryHandler {
	return &categoryHandler{service}
}

/*
GET LIST

Tangkap Parameter di Handler
kirim handle ke service
Service menentukan repository mana yg akan d Call(Get All atau ByUser)
Repository : GetAll & GetByUser
DB
*/
func (h *categoryHandler) GetCategory(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	categories, err := h.service.GetCategory(userID)
	if err != nil {
		responsError := helper.APIResponse("Get Category Failed #CG0019", http.StatusBadRequest, "fail", nil)
		c.JSON(http.StatusBadRequest, responsError)
		return
	}

	responsSuccess := helper.APIResponse("Get Category Success", http.StatusOK, "success", category.FormatCategories(categories))
	c.JSON(http.StatusOK, responsSuccess)

}

/*
GET DETAIL
*/
func (h *categoryHandler) GetCategoryDetail(c *gin.Context) {
	var input category.CategoryUserInput

	err := c.ShouldBindUri(&input)

	if err != nil {
		responsError := helper.APIResponse("Get Category Detail Failed #CGD0871", http.StatusBadRequest, "fail", nil)
		c.JSON(http.StatusBadRequest, responsError)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User).ID

	categoryDetail, err := h.service.GetDetailCategory(input, currentUser)

	if err != nil {
		responsError := helper.APIResponse("Get Category Detail Failed #CGD0872", http.StatusUnauthorized, "fail", nil)
		c.JSON(http.StatusUnauthorized, responsError)
		return
	}

	response := helper.APIResponse("Get Category Detail Success", http.StatusOK, "success", category.FormatCategoryDetail(categoryDetail))
	c.JSON(http.StatusOK, response)
	return
}

func (h *categoryHandler) CreateCategory(c *gin.Context) {
	var input category.CreateCategoryInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}

		responsError := helper.APIResponse("Create Category Failed #CRC090", http.StatusUnprocessableEntity, "fail", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, responsError)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	input.User = currentUser

	newCategory, err := h.service.CreateCategory(input)
	if err != nil {
		responsError := helper.APIResponse("Create Category Failed #CRC097", http.StatusBadGateway, "fail", nil)
		c.JSON(http.StatusBadGateway, responsError)
		return
	}

	response := helper.APIResponse("Create Category Success", http.StatusOK, "success", category.FormatCategory(newCategory))
	c.JSON(http.StatusOK, response)

}

/*
Update Category

*/

func (h *categoryHandler) UpdateCategory(c *gin.Context) {
	var inputID category.CategoryUserInput

	err := c.ShouldBindUri(&inputID)

	if err != nil {
		responsError := helper.APIResponse("Update Category Failed #CGU0081", http.StatusBadRequest, "fail", nil)
		c.JSON(http.StatusBadRequest, responsError)
		return
	}

	var inputData category.CreateCategoryInput

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}

		responsError := helper.APIResponse("Update Category Failed #CGU0085", http.StatusUnprocessableEntity, "fail", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, responsError)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	inputData.User = currentUser

	updatedCategory, err := h.service.UpdateCategory(inputID, inputData)
	if err != nil {
		responsError := helper.APIResponse("Update Category Failed #CGU0093", http.StatusUnauthorized, "fail", nil)
		c.JSON(http.StatusUnauthorized, responsError)
		return
	}

	response := helper.APIResponse("Update Category Success", http.StatusOK, "success", category.FormatCategory(updatedCategory))
	c.JSON(http.StatusOK, response)
}

func (h *categoryHandler) DeleteCategory(c *gin.Context) {
	var input category.CategoryUserInput

	err := c.ShouldBindUri(&input)

	if err != nil {
		responsError := helper.APIResponse("Delete Category Failed #CHT091", http.StatusBadRequest, "fail", nil)
		c.JSON(http.StatusBadRequest, responsError)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User).ID

	_, err = h.service.DeleteCategory(input, currentUser)
	if err != nil {
		responsError := helper.APIResponse("Delete Category Failed #CHT017", http.StatusUnauthorized, "fail", nil)
		c.JSON(http.StatusUnauthorized, responsError)
		return
	}

	response := helper.APIResponse("Delete Category Success", http.StatusOK, "success", true)
	c.JSON(http.StatusOK, response)
	return
}
