package user

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-playground/validator"
)

// func ValidateUser(user User) error {
// 	validate := validator.New()

// 	err := validate.Struct(user)
// 	if err != nil {
// 		validationErrors := err.(validator.ValidationErrors)
// 		errorMessages := make([]string, len(validationErrors))

// 		for i, validationErr := range validationErrors {
// 			fieldName := validationErr.Field()
// 			switch fieldName {
// 			case "Email":
// 				errorMessages[i] = "Invalid Email"
// 			case "Username":
// 				errorMessages[i] = "Invalid Username, Minimum 8 letters or Maximum 24 letters required"
// 			case "FirstName":
// 				errorMessages[i] = "Invalid Firstname, Minimum 4 letters or Maximum 10 letters required"
// 			case "LastName":
// 				errorMessages[i] = "Invalid Lastname, Minimum 4 letters or Maximum 10 letters required"
// 			case "Password":
// 				errorMessages[i] = "Invalid password, Minimum 6 letters or Maximum 12 letters required"
// 			case "PhoneNumber":
// 				errorMessages[i] = "Invalid Phone Number"
// 			default:
// 				errorMessages[i] = "Validation failed"
// 			}
// 		}

//			return fmt.Errorf(strings.Join(errorMessages, ", "))
//		}
//		return nil
//	}

func ValidateUser(user User) error {
	validate := validator.New()

	// Register custom date validation (YYYY-MM-DD)
	validate.RegisterValidation("dateformat", func(fl validator.FieldLevel) bool {
		date := fl.Field().String()
		if date == "" {
			return true
		}
		matched, _ := regexp.MatchString(`^\d{4}-\d{2}-\d{2}$`, date)
		return matched
	})

	err := validate.Struct(user)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := make([]string, len(validationErrors))

		for i, validationErr := range validationErrors {
			fieldName := validationErr.Field()
			switch fieldName {

			case "Email":
				errorMessages[i] = "Invalid Email"

			case "Username":
				errorMessages[i] = "Invalid Username, Minimum 8 letters or Maximum 24 letters required"

			case "FirstName":
				errorMessages[i] = "Invalid Firstname, Minimum 4 letters or Maximum 10 letters required"

			case "LastName":
				errorMessages[i] = "Invalid Lastname, Minimum 4 letters or Maximum 10 letters required"

			case "Password":
				errorMessages[i] = "Invalid password, Minimum 6 letters or Maximum 12 letters required"

			case "PhoneNumber":
				errorMessages[i] = "Invalid Phone Number"

			case "DateOfBirth":
				errorMessages[i] = "Invalid Date of Birth, format should be YYYY-MM-DD"

			default:
				errorMessages[i] = "Validation failed"
			}
		}

		return fmt.Errorf(strings.Join(errorMessages, ", "))
	}

	return nil
}

// func ValidateUserProfile(profile UserProfileDetails) error {
// 	validate := validator.New()

// 	// Reuse the same date format validation
// 	validate.RegisterValidation("dateformat", func(fl validator.FieldLevel) bool {
// 		date := fl.Field().String()
// 		if date == "" {
// 			return true
// 		}
// 		matched, _ := regexp.MatchString(`^\d{4}-\d{2}-\d{2}$`, date)
// 		return matched
// 	})

// 	err := validate.Struct(profile)
// 	if err != nil {
// 		validationErrors := err.(validator.ValidationErrors)
// 		errorMessages := make([]string, len(validationErrors))

// 		for i, validationErr := range validationErrors {
// 			fieldName := validationErr.Field()
// 			switch fieldName {

// 			case "Email":
// 				errorMessages[i] = "Invalid Email"

// 			case "Username":
// 				errorMessages[i] = "Invalid Username, Minimum 8 or Maximum 24 characters"

// 			case "FirstName":
// 				errorMessages[i] = "Invalid Firstname, Minimum 4 or Maximum 10 characters"

// 			case "LastName":
// 				errorMessages[i] = "Invalid Lastname, Minimum 4 or Maximum 10 characters"

// 			case "PhoneNumber":
// 				errorMessages[i] = "Invalid Phone Number"

// 			case "DateOfBirth":
// 				errorMessages[i] = "Invalid Date of Birth (YYYY-MM-DD)"

// 			default:
// 				errorMessages[i] = "Validation failed"
// 			}
// 		}

//			return fmt.Errorf(strings.Join(errorMessages, ", "))
//		}
//		return nil
//	}

// func ValidateUpdate(update UserProfileDetails) error {
// 	validate := validator.New()

// 	validate.RegisterValidation("dateformat", func(fl validator.FieldLevel) bool {
// 		date := fl.Field().String()
// 		if date == "" {
// 			return true
// 		}
// 		matched, _ := regexp.MatchString(`^\d{4}-\d{2}-\d{2}$`, date)
// 		return matched
// 	})

// 	if err := validate.Struct(update); err != nil {
// 		validationErrors := err.(validator.ValidationErrors)
// 		errorMessages := make([]string, len(validationErrors))

// 		for i, v := range validationErrors {
// 			switch v.Field() {
// 			case "Email":
// 				errorMessages[i] = "Invalid email format"
// 			case "Username":
// 				errorMessages[i] = "Username must be 8–24 characters"
// 			case "FirstName":
// 				errorMessages[i] = "Firstname must be 4–10 characters"
// 			case "LastName":
// 				errorMessages[i] = "Lastname must be 4–10 characters"
// 			case "PhoneNumber":
// 				errorMessages[i] = "Phone must be 10 digits"
// 			case "DateOfBirth":
// 				errorMessages[i] = "DOB must be YYYY-MM-DD"
// 			default:
// 				errorMessages[i] = "Validation failed"
// 			}
// 		}

// 		return fmt.Errorf(strings.Join(errorMessages, ", "))
// 	}

// 	return nil
// }

func ValidateUpdate(update UserProfileDetails) error {
	validate := validator.New()

	validate.RegisterValidation("dateformat", func(fl validator.FieldLevel) bool {
		date := fl.Field().String()
		if date == "" {
			return true
		}
		matched, _ := regexp.MatchString(`^\d{4}-\d{2}-\d{2}$`, date)
		return matched
	})

	if err := validate.Struct(update); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := make([]string, len(validationErrors))

		for i, v := range validationErrors {
			switch v.Field() {
			case "Email":
				errorMessages[i] = "Invalid email format"
			case "Username":
				errorMessages[i] = "Username must be 8–24 characters"
			case "FirstName":
				errorMessages[i] = "Firstname must be 4–10 characters"
			case "LastName":
				errorMessages[i] = "Lastname must be 4–10 characters"
			case "PhoneNumber":
				errorMessages[i] = "Phone must be 10 digits"
			case "DateOfBirth":
				errorMessages[i] = "DOB must be YYYY-MM-DD"
			default:
				errorMessages[i] = "Validation failed"
			}
		}

		return fmt.Errorf(strings.Join(errorMessages, ", "))
	}

	return nil
}
