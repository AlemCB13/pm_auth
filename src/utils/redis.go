package utils

import (
	"context"
	"errors"
	"pm_auth/src/models"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "pm_auth_redis:6379", // Dirección de Redis
		Password: "",                   // Sin contraseña
		DB:       0,                    // Base de datos por defecto
	})
}

func SaveUser(user models.User) error {
	// Guardar el usuario en Redis
	err := rdb.Set(ctx, user.Username, user.Password, 0).Err()
	return err
}

func VerifyCredentials(credentials models.Credentials) (string, error) {
	// Verificar las credenciales en Redis
	val, err := rdb.Get(ctx, credentials.Username).Result()
	if err != nil {
		return "", err
	}

	if val != credentials.Password {
		return "", errors.New("credenciales inválidas")
	}

	// Generar un token JWT (implementar esta función)
	token, err := GenerateToken(credentials.Username)
	return token, err
}

func InvalidateToken(token string) error {
	// Invalidar el token en Redis
	err := rdb.Set(ctx, token, "invalid", 0).Err()
	return err
}
