package helper

import (
	"golang.org/x/crypto/bcrypt"
)

/* const (
	saltSize = 32
	N        = 16384
	r        = 8
	p        = 1
	keyLen   = 32
)

const dummyPwd = "aBc1R42e!@kQ9.2f"

var dummySalt = []byte{23, 134, 65, 223, 98, 34, 75, 134, 203, 10, 46, 32, 164, 203, 27, 99,
	213, 41, 19, 4, 9, 174, 203, 109, 59, 77, 187, 219, 43, 19, 78, 90}

type PasswordHash struct {
	Hash []byte
	Salt []byte
}

func GeneratePasswordHash(password string, salt []byte) (*PasswordHash, error) {
	if salt == nil {
		salt = make([]byte, saltSize)
		_, err := io.ReadFull(rand.Reader, salt)
		if err != nil {
			log.Printf("salt generation failed: %v", err)
			return nil, err
		}
	}
	hash, err := scrypt.Key([]byte(password), salt, N, r, p, keyLen)
	if err != nil {
		return nil, err
	}
	ph := PasswordHash{Hash: hash, Salt: salt}
	return &ph, nil
}

func DoDummyPasswordHash() error {
	_, err := GeneratePasswordHash(dummyPwd, dummySalt)
	return err
} */

func GeneratePasswordHash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.MaxCost)
}

func VerifyPassword(hashedPassword, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}
