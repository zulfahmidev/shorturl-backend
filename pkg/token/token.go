package token

import (
	"crypto/rand"
	"math/big"

	"github.com/zulfahmidev/shorturl/pkg/database"
)

const tokenLength = 10
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateRandomToken() string {
	b := make([]byte, tokenLength)
	for i := range b {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		b[i] = charset[num.Int64()]
	}
	return string(b)
}

func GenerateUniqueToken() (string, error) {
	db := database.GetDB()
	for {
		token := GenerateRandomToken()
		var exists bool
		db.Exec("SELECT EXISTS(SELECT 1 FROM urls WHERE token=$1)", token).Scan(&exists)
		if !exists {
			return token, nil
		}
	}
}