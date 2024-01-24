package helpers

import (
	"errors"
	"fmt"
	"strings"
)

// BodyRequest is a struct to hold the request body.
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

func ValidateFields(bodyRequest *BodyRequest) error {
	// Required fields
	requiredFields := map[string]string{
		"Name":    bodyRequest.Name,
		"Phone":   bodyRequest.Phone,
		"Email":   bodyRequest.Email,
		"Address": bodyRequest.Address,
	}

	// Validate required fields
	for fieldName, fieldValue := range requiredFields {
		// Remove extra spaces
		cleanedFieldValue := strings.Join(strings.Fields(fieldValue), " ")

		if cleanedFieldValue == "" {
			return fmt.Errorf("%s is required", fieldName)
		}

		// Check if email field contains an @ symbol
		if fieldName == "Email" && !strings.Contains(cleanedFieldValue, "@") {
			return errors.New("Email must contain an @ symbol")
		}

		// Assign cleaned value back to the struct
		switch fieldName {
		case "Name":
			bodyRequest.Name = cleanedFieldValue
		case "Phone":
			bodyRequest.Phone = cleanedFieldValue
		case "Email":
			bodyRequest.Email = cleanedFieldValue
		case "Address":
			bodyRequest.Address = cleanedFieldValue
		}
	}

	return nil
}
