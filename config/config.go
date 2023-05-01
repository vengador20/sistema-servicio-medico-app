package config

import (
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var (
	once          sync.Once
	env           string
	privateKeyJwt string
)

// obtener variable de entorno
func GetEnviroment() (string, error) {
	var err error
	once.Do(func() {
		err = godotenv.Load()
		env = os.Getenv("MONGO_URL")
		privateKeyJwt = os.Getenv("PRIVATEKEY_JWT")
	})

	if err != nil {
		return "", err
	}

	return env, nil
}

func PrivateKeyJwt() string {
	return privateKeyJwt
}
