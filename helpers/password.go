package helpers

import "golang.org/x/crypto/bcrypt"

// HashPassword generates a hashed password from given input
func HashPassword(password string) (string, error) {
	passwordByte := []byte(password)
	defer clearMemory(passwordByte)
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordByte, bcrypt.DefaultCost)
	return string(hashedPassword), err
}

// ComparePassword compares the hashed password and the password
func ComparePassword(hashedPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err
}

func clearMemory(b []byte) {
	for i := 0; i < len(b); i++ {
		b[i] = 0
	}
}
