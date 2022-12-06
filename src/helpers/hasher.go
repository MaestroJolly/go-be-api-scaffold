package helpers

import "golang.org/x/crypto/bcrypt"

// function password encryption
func encryptor(password string) (*string, error) {
	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	stringifiedHashedPassword := string(hashedPassword)
	return &stringifiedHashedPassword, nil
}
