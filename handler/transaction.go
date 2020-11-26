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

func (h *transactionHandler) GetTransaction(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	transaction, err := h.service.GetTransaction(userID)
	if err != nil {
		responsError := helper.APIResponse("Get Transaction Failed #CG0019", http.StatusBadRequest, "fail", nil)
		c.JSON(http.StatusBadRequest, responsError)
		return
	}

	responsSuccess := helper.APIResponse("Get Transaction Success", http.StatusOK, "success", transactions.FormatTransactions(transaction))
	c.JSON(http.StatusOK, responsSuccess)

}

func (h *transactionHandler) GetTransactionByDates(c *gin.Context) {
	var input transactions.TransactionFilterDate

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}

		responsError := helper.APIResponse("Get List Transaction Failed #CRC090", http.StatusUnprocessableEntity, "fail", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, responsError)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	transaction, err := h.service.GetTransactionByDate(userID, input.StartDate, input.EndDate)
	if err != nil {
		responsError := helper.APIResponse("Get Transaction Failed #CG0019", http.StatusBadRequest, "fail", nil)
		c.JSON(http.StatusBadRequest, responsError)
		return
	}

	responsSuccess := helper.APIResponse("Get Transaction Success", http.StatusOK, "success", transactions.FormatTransactions(transaction))
	c.JSON(http.StatusOK, responsSuccess)

}

func (h *transactionHandler) CreateTransaction(c *gin.Context) {
	var input transactions.TransactionUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}

		responsError := helper.APIResponse("Create Transaction Failed #CRC090", http.StatusUnprocessableEntity, "fail", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, responsError)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	// input.User = currentUser
	checkcategory, err := h.service.CheckTypeCategory(currentUser.ID, input.CategoryID)

	if err != nil {
		if err != nil {
			responsError := helper.APIResponse("Create Transaction Failed #CRC097", http.StatusUnauthorized, "fail", nil)
			c.JSON(http.StatusUnauthorized, responsError)
			return
		}
	}
	nominal := 0
	if checkcategory.Type == "Pemasukan" {
		nominal = currentUser.Balance + input.Nominal
		_, err := h.service.UpdateUser(currentUser.ID, nominal)

		if err != nil {
			responsError := helper.APIResponse("Create Transaction Failed #CRC097", http.StatusBadGateway, "fail", nil)
			c.JSON(http.StatusBadGateway, responsError)
			return
		}
	} else {
		nominal = currentUser.Balance - input.Nominal
		_, err := h.service.UpdateUser(currentUser.ID, nominal)

		if err != nil {
			responsError := helper.APIResponse("Create Transaction Failed #CRC097", http.StatusBadGateway, "fail", nil)
			c.JSON(http.StatusBadGateway, responsError)
			return
		}
	}
	input.User = currentUser
	newTransaction, err := h.service.SaveTransaction(input)
	if err != nil {
		responsError := helper.APIResponse("Create Transaction Failed #CRC097", http.StatusBadGateway, "fail", nil)
		c.JSON(http.StatusBadGateway, responsError)
		return
	}

	response := helper.APIResponse("Create Transaction Success", http.StatusOK, "success", transactions.FormatTransactionCreate(newTransaction))
	c.JSON(http.StatusOK, response)

}

func (h *transactionHandler) GetTransactionDetail(c *gin.Context) {
	var input transactions.IDUserInput

	err := c.ShouldBindUri(&input)

	if err != nil {
		responsError := helper.APIResponse("Get Transaction Detail Failed #CGD0871", http.StatusBadRequest, "fail", nil)
		c.JSON(http.StatusBadRequest, responsError)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User).ID

	transactionDetail, err := h.service.GetDetailTransaction(input, currentUser)
	if err != nil {
		responsError := helper.APIResponse("Get Transaction Detail Failed #CGD0872", http.StatusUnauthorized, "fail", nil)
		c.JSON(http.StatusUnauthorized, responsError)
		return
	}

	response := helper.APIResponse("Get Transaction Detail Success", http.StatusOK, "success", transactions.FormatTransactionDetail(transactionDetail))
	c.JSON(http.StatusOK, response)
	return
}

func (h *transactionHandler) DeleteTransaction(c *gin.Context) {
	var input transactions.IDUserInput

	err := c.ShouldBindUri(&input)

	if err != nil {
		responsError := helper.APIResponse("Delete Transaction Failed #TRX0018", http.StatusBadRequest, "fail", nil)
		c.JSON(http.StatusBadRequest, responsError)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User).ID

	_, err = h.service.DeleteTransaction(input, currentUser)
	if err != nil {
		responsError := helper.APIResponse("Delete Transaction Failed #TRX01891", http.StatusUnauthorized, "fail", nil)
		c.JSON(http.StatusUnauthorized, responsError)
		return
	}

	response := helper.APIResponse("Delete Transaction Success", http.StatusOK, "success", true)
	c.JSON(http.StatusOK, response)
	return
}
