package handler

import (
	"TWallet/helper"
	"TWallet/transactions"
	"TWallet/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	service transactions.Service
	// serviceCategory category.Service
}

func NewTransactionHandler(service transactions.Service) *transactionHandler {
	return &transactionHandler{service}
}

func (h *transactionHandler) CreateTransaction(c *gin.Context) {
	var input transactions.TransactionUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}

		responsError := helper.APIResponse("Create Category Failed #CRC090", http.StatusUnprocessableEntity, "fail", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, responsError)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	// input.User = currentUser
	checkcategory, err := h.service.CheckTypeCategory(currentUser.ID, input.CategoryID)

	if err != nil {
		if err != nil {
			responsError := helper.APIResponse("Create Campaign Failed #CRC097", http.StatusUnauthorized, "fail", nil)
			c.JSON(http.StatusUnauthorized, responsError)
			return
		}
	}
	nominal := 0
	if checkcategory.Type == "Pemasukan" {
		nominal = currentUser.Balance + input.Nominal
		_, err := h.service.UpdateUser(currentUser.ID, nominal)

		if err != nil {
			responsError := helper.APIResponse("Create Campaign Failed #CRC097", http.StatusBadGateway, "fail", nil)
			c.JSON(http.StatusBadGateway, responsError)
			return
		}
	} else {
		nominal = currentUser.Balance - input.Nominal
		_, err := h.service.UpdateUser(currentUser.ID, nominal)

		if err != nil {
			responsError := helper.APIResponse("Create Campaign Failed #CRC097", http.StatusBadGateway, "fail", nil)
			c.JSON(http.StatusBadGateway, responsError)
			return
		}
	}
	input.User = currentUser
	newCampaign, err := h.service.SaveTransaction(input)
	if err != nil {
		responsError := helper.APIResponse("Create Campaign Failed #CRC097", http.StatusBadGateway, "fail", nil)
		c.JSON(http.StatusBadGateway, responsError)
		return
	}

	response := helper.APIResponse("Create Campaign Success", http.StatusOK, "success", newCampaign)
	c.JSON(http.StatusOK, response)

}
