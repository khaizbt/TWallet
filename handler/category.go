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

	campaigns, err := h.service.GetCategory(userID)
	if err != nil {
		responsError := helper.APIResponse("Get Campaign Failed #CG0019", http.StatusBadRequest, "fail", nil)
		c.JSON(http.StatusBadRequest, responsError)
		return
	}

	responsSuccess := helper.APIResponse("Get Campaign Success", http.StatusOK, "success", category.FormatCategories(campaigns))
	c.JSON(http.StatusOK, responsSuccess)

}

/*
GET DETAIL

url /api/v1/campaign/:id-campaign
Mapping id yang di URL ke struct input lalu parsing ke Service kemudian panggi formater
Service : Struct Input dengan menangkap id dari url manggil repo
Repository : get campaign by id
DB
*/
func (h *categoryHandler) GetCategoryDetail(c *gin.Context) {
	var input category.CategoryUserInput

	err := c.ShouldBindUri(&input)

	if err != nil {
		responsError := helper.APIResponse("Get Campaign Detail Failed #CGD0871", http.StatusBadRequest, "fail", nil)
		c.JSON(http.StatusBadRequest, responsError)
		return
	}

	campaignDetail, err := h.service.GetDetailCategory(input)

	if err != nil {
		responsError := helper.APIResponse("Get Campaign Detail Failed #CGD0872", http.StatusBadRequest, "fail", nil)
		c.JSON(http.StatusBadRequest, responsError)
		return
	}

	response := helper.APIResponse("Get Campaign Detail Success", http.StatusOK, "success", category.FormatCampaignDetail(campaignDetail))
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

	newCampaign, err := h.service.CreateCategory(input)
	if err != nil {
		responsError := helper.APIResponse("Create Campaign Failed #CRC097", http.StatusBadGateway, "fail", nil)
		c.JSON(http.StatusBadGateway, responsError)
		return
	}

	response := helper.APIResponse("Create Campaign Success", http.StatusOK, "success", category.FormatCategory(newCampaign))
	c.JSON(http.StatusOK, response)

}

/*
Update Campaign
1. User Melakukan Input
2. Handler
3. Mapping dari input dan URI ke struct
4. Service
5. Repository
*/

func (h *categoryHandler) UpdateCategory(c *gin.Context) {
	var inputID category.CategoryUserInput

	err := c.ShouldBindUri(&inputID)

	if err != nil {
		responsError := helper.APIResponse("Update Campaign Failed #CGU0081", http.StatusBadRequest, "fail", nil)
		c.JSON(http.StatusBadRequest, responsError)
		return
	}

	var inputData category.CreateCategoryInput

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}

		responsError := helper.APIResponse("Update Campaign Failed #CGU0085", http.StatusUnprocessableEntity, "fail", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, responsError)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	inputData.User = currentUser

	updatedCampaign, err := h.service.UpdateCampaign(inputID, inputData)
	if err != nil {
		responsError := helper.APIResponse("Update Campaign Failed #CGU0093", http.StatusBadRequest, "fail", nil)
		c.JSON(http.StatusBadRequest, responsError)
		return
	}

	response := helper.APIResponse("Update Campaign Success", http.StatusOK, "success", category.FormatCategory(updatedCampaign))
	c.JSON(http.StatusOK, response)
}
