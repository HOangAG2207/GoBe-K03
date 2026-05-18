package service

import (
	"crypto/rand"
	"errors"
	"math/big"
)

// defaultCharset defines the set of characters used for password generation.
// It includes:
// - lowercase letters (a-z)
// - uppercase letters (A-Z)
// - digits (0-9)
const defaultCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// GenPassword defines the contract for password generation

//go:generate mockery --name GenPassword --filename gen_password_mock.go --output ./mocks
type GenPassword interface {
	// GeneratePassword generates a random password with the given length
	// Input:
	//   - length: desired password length
	// Output:
	//   - string: generated password
	//   - error: returned if generation fails
	GeneratePassword(length int) (string, error)
}

// genPasswordService is the concrete implementation of GenPassword
type genPasswordService struct {
	// charset is the pool of characters used to build the password
	charset string
}

// NewGenPassword initializes the password generator service
// Return:
//   - GenPassword: interface for external usage
func NewGenPassword() GenPassword {
	return &genPasswordService{
		charset: defaultCharset,
	}
}

// GeneratePassword generates a random password
func (s *genPasswordService) GeneratePassword(length int) (string, error) {

	// ===== 1. Validate input =====

	// length must be greater than 0
	if length <= 0 {
		return "", errors.New("length must be greater than 0")
	}
	// ===== 2. Prepare data =====

	// max defines the upper bound for random index generation
	// rand.Int generates a number in range [0, max)
	max := big.NewInt(int64(len(s.charset)))

	// allocate byte slice for password
	password := make([]byte, length)

	// ===== 3. Generate characters =====

	for i := range length {

		// generate a secure random index
		idx, err := rand.Int(rand.Reader, max)
		if err != nil {
			// return immediately if crypto random fails
			return "", err
		}

		// map index to character in charset
		password[i] = s.charset[idx.Int64()]
	}

	// ===== 4. Convert & return =====

	// convert byte slice to string and return
	return string(password), nil
}
