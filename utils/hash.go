package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	// Hash the password
	bcryptPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bcryptPassword)
	
}

func ComparePassword(hashedPassword, password string) bool {
	// Compare the password with the hash
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}