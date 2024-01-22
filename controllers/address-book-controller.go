package controllers

import (
	"address-book-backend/helpers"
	"address-book-backend/initializers"
	"address-book-backend/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateAddressBook is a function to create a new address book entry in the database.
func CreateAddressBook(c *gin.Context) {
	// Define a struct to hold the request body. The `binding:"required"` tag means
	// that the field must be present in the request body.
	var bodyRequest struct {
		Name     string `json:"name" binding:"required"`
		Phone    string `json:"phone" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Address  string `json:"address" binding:"required"`
		Company  string `json:"company"`
		JobTitle string `json:"job_title"`
		Website  string `json:"website"`
		Note     string `json:"note"`
	}

	// Bind the request body to the bodyRequest struct.
	c.Bind(&bodyRequest)

	// Validate the required fields. If a field is empty, respond with an error message
	// and a status code of 400 (Bad Request).
	fields := map[string]string{
		"Name":    bodyRequest.Name,
		"Phone":   bodyRequest.Phone,
		"Email":   bodyRequest.Email,
		"Address": bodyRequest.Address,
	}
	for fieldName, fieldValue := range fields {
		if fieldValue == "" {
			helpers.RespondWithError(c, http.StatusBadRequest, fmt.Sprintf("%s is required", fieldName), "01")
			return
		}
	}

	// Create a new AddressBook model with the data from the request body.
	addreesBook := models.AddressBook{
		Name:     bodyRequest.Name,
		Phone:    bodyRequest.Phone,
		Email:    bodyRequest.Email,
		Company:  bodyRequest.Company,
		JobTitle: bodyRequest.JobTitle,
		Website:  bodyRequest.Website,
		Address:  bodyRequest.Address,
		Note:     bodyRequest.Note,
	}

	// Use GORM's Create method to insert the new address book entry into the database.
	result := initializers.DB.Create(&addreesBook)

	// If there's an error during the Create operation, respond with an error message
	// and a status code of 400 (Bad Request).
	if result.Error != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, result.Error, "01")
		return
	}

	// If there's no error, respond with a success message, a status code of 201 (Created),
	// and the new address book entry.
	helpers.RespondWithSuccess(c, http.StatusCreated, "Address book created successfully", "00", addreesBook)
}

// GetAddressBooks is a function to retrieve all address books from the database.
func GetAddressBooks(c *gin.Context) {
	// Declare a slice of AddressBook models.
	var addressBook []models.AddressBook

	// Use GORM's Find method to retrieve all records from the address_books table
	// and assign them to the addressBook slice.
	result := initializers.DB.Find(&addressBook)

	// If there's an error during the Find operation, respond with an error message
	// and a status code of 400 (Bad Request).
	if result.Error != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, result.Error, "01")
		return
	}

	// If there's no error, respond with a success message, a status code of 200 (OK),
	// and the addressBook slice containing the retrieved records.
	helpers.RespondWithSuccess(c, http.StatusOK, "Address books retrieved successfully", "00", addressBook)
}

// GetAddressBook is a function to retrieve a single address book entry from the database.
func GetAddressBook(c *gin.Context) {
	// Declare an AddressBook model.
	var addressBook models.AddressBook

	// Use GORM's First method to retrieve the first record from the address_books table
	// and assign it to the addressBook variable.
	result := initializers.DB.First(&addressBook, c.Param("id"))

	// If there's an error during the First operation, respond with an error message
	// and a status code of 400 (Bad Request).
	if result.Error != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, result.Error, "01")
		return
	}

	// If there's no error, respond with a success message, a status code of 200 (OK),
	// and the addressBook variable containing the retrieved record.
	helpers.RespondWithSuccess(c, http.StatusOK, "Address book retrieved successfully", "00", addressBook)
}

// UpdateAddressBook is a function to update a single address book entry in the database.
func UpdateAddressBook(c *gin.Context) {
	// Declare an AddressBook model.
	var addressBook models.AddressBook

	// Use GORM's First method to retrieve the first record from the address_books table
	// and assign it to the addressBook variable.
	result := initializers.DB.First(&addressBook, c.Param("id"))

	// If there's an error during the First operation, respond with an error message
	// and a status code of 400 (Bad Request).
	if result.Error != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, result.Error, "01")
		return
	}

	// Define a struct to hold the request body. This time, the user can update just one or more
	// fields, so the `binding:"required"` tag is removed.

	var bodyRequest struct {
		Name     string `json:"name"`
		Phone    string `json:"phone"`
		Email    string `json:"email"`
		Address  string `json:"address"`
		Company  string `json:"company"`
		JobTitle string `json:"job_title"`
		Website  string `json:"website"`
		Note     string `json:"note"`
	}

	// Bind the request body to the bodyRequest struct.
	c.Bind(&bodyRequest)

	// Update the addressBook variable with the data from the request body.
	addressBook.Name = bodyRequest.Name
	addressBook.Phone = bodyRequest.Phone
	addressBook.Email = bodyRequest.Email
	addressBook.Address = bodyRequest.Address
	addressBook.Company = bodyRequest.Company
	addressBook.JobTitle = bodyRequest.JobTitle
	addressBook.Website = bodyRequest.Website
	addressBook.Note = bodyRequest.Note

	// Use GORM's Update method to update the address book entry in the database.
	result = initializers.DB.Model(&addressBook).Updates(addressBook)

	// If there's an error during the Save operation, respond with an error message
	// and a status code of 400 (Bad Request).
	if result.Error != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, result.Error, "01")
		return
	}

	// If there's no error, respond with a success message, a status code of 200 (OK),
	// and the addressBook variable containing the updated record.
	helpers.RespondWithSuccess(c, http.StatusOK, "Address book updated successfully", "00")
}

// DeleteAddressBook is a function to delete a single address book entry from the database.
func DeleteAddressBook(c *gin.Context) {
	// Declare an AddressBook model.
	var addressBook models.AddressBook

	// Use GORM's First method to retrieve the first record from the address_books table
	// and assign it to the addressBook variable.
	result := initializers.DB.First(&addressBook, c.Param("id"))

	// If there's an error during the First operation, respond with an error message
	// and a status code of 400 (Bad Request).
	if result.Error != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, result.Error, "01")
		return
	}

	// Use GORM's Delete method to delete the address book entry from the database.
	result = initializers.DB.Delete(&addressBook)

	// If there's an error during the Delete operation, respond with an error message
	// and a status code of 400 (Bad Request).

	if result.Error != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, result.Error, "01")
		return
	}

	// If there's no error, respond with a success message, a status code of 200 (OK),
	// and the addressBook variable containing the deleted record.
	helpers.RespondWithSuccess(c, http.StatusOK, "Address book deleted successfully", "00")

}
