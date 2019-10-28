package models

import "crypto/sha1"

//User defines the user mongoDB schema
type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
}

//UserCredentials defines the login api schema
type UserCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//Validate ensures that user is in a correct format
func (user *User) Validate() []string {
	var errorData []string = []string{}
	if user.FirstName == "" {
		errorData = append(errorData, "user.firstname: field is required")
	}
	if user.LastName == "" {
		errorData = append(errorData, "user.lastname: field is required")
	}
	if user.Email == "" {
		errorData = append(errorData, "user.email: field is required")
	}
	if user.Password == "" {
		errorData = append(errorData, "user.password: field is required")
	}
	if user.Role == "" {
		errorData = append(errorData, "user.role: field is required")
	}
	return errorData
}

//Validate checks login request payload
func (userCredentials *UserCredentials) Validate() []string {
	var errorData []string = []string{}
	if userCredentials.Email == "" {
		errorData = append(errorData, "email: field is required")
	}
	if userCredentials.Password == "" {
		errorData = append(errorData, "password: field is required")
	}
	return errorData
}

//IsPasswordCorrect hashes the password and checks if it matches the password hash
func (user *User) IsPasswordCorrect(password string) bool {
	hasher := sha1.New()
	hasher.Write([]byte(password))
	res := hasher.Sum(nil)
	if string(res) == user.Password {
		return true
	}
	return false
}
