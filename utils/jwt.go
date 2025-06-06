package utils

import (
    "time"
    "github.com/golang-jwt/jwt/v5"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "fmt"
)

var claveSecreta = []byte("clave_secreta")

func CrearToken(_id primitive.ObjectID, nombre string) (string, error) {
    claims := jwt.MapClaims{
        "userID": _id.Hex(), // Convertir ObjectID a string
        "user":   nombre,
        "exp": time.Now().Add(10 * time.Minute).Unix(),
		"iat": time.Now().Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(claveSecreta)
}

func ValidarToken(tokenString string) (*jwt.Token, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("método de firma inesperado: %v", token.Header["alg"])
        }
        return claveSecreta, nil
    })

    if err != nil {
        return nil, err
    }

    return token, nil
}
