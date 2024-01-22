package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BodyRequest struct {
	Name     string `json:"name" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Company  string `json:"company"`
	JobTitle string `json:"job_title"`
	Website  string `json:"website"`
	Note     string `json:"note"`
}

func ValidateBodyRequest(c *gin.Context, bodyRequest BodyRequest) {
	fields := map[string]string{
		"Name":    bodyRequest.Name,
		"Phone":   bodyRequest.Phone,
		"Email":   bodyRequest.Email,
		"Address": bodyRequest.Address,
	}

	for fieldName, fieldValue := range fields {
		if fieldValue == "" {
			RespondWithError(c, http.StatusBadRequest, fieldName+" is required", "01")
			return
		}
	}
}
