package auth

import "github.com/alexedwards/argon2id"

func HashPassword(password string) (string, error) {
	haspPass, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return "", err
	}

	return haspPass, nil
}

func CheckPasswordHash(password string, hash string) (bool, error) {
	isPassword, err := argon2id.ComparePasswordAndHash(password, hash)
	if err != nil {
		return false, err
	}

	return isPassword, nil
}
