package hash

import "golang.org/x/crypto/bcrypt"

func Hash(text string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(text), 10)
	return string(bytes), err
}

func Verify(text string, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(text))
}
