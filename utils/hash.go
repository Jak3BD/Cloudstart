package utils

import "github.com/ncw/pwhash/sha512_crypt"

func HashPassword(password string) string {
	salt := sha512_crypt.GenerateSalt(16, sha512_crypt.RoundsDefault)
	hashedPassword := sha512_crypt.Crypt(password, salt)

	return hashedPassword
}
