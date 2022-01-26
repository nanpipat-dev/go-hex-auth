package security

import (
	"errors"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"encoding/base64"
)

var (
	JwtSecretKey     = []byte("secret")
	JwtSigningMethod = jwt.SigningMethodHS256.Name
)

// Initialize victor, which is the random bytes

func NewToken(userId string) (string, error) {
	fmt.Println(JwtSecretKey)
	claims :=
		jwt.StandardClaims{
			Id:        userId,
			ExpiresAt: time.Now().Add((time.Minute * 15)).Unix(),
		}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtSecretKey)
}

func RefreshToken(userId string) (string, error) {
	claims :=
		jwt.StandardClaims{
			Id:        userId,
			ExpiresAt: time.Now().Add((time.Minute * 500)).Unix(),
		}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	getTk, err := token.SignedString(JwtSecretKey)
	if err != nil {
		return "", err
	}

	// swEnc := base64.RawStdEncoding.EncodeToString([]byte(getTk))
	// swEnc2, err := base64.RawStdEncoding.DecodeString(swEnc)
	// if err != nil {
	// 	return "", err
	// }

	// fmt.Println(string(swEnc2), "decode")

	refresh := encrypt(getTk)

	return refresh, nil
}

func ParseToken(tokenString string) (*jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtSecretKey, nil
	})

	if err != nil {
		v, _ := err.(*jwt.ValidationError)
		fmt.Println(v, "err")
		return nil, err
	}

	// var ok bool

	claims := token.Claims.(*jwt.StandardClaims)

	if !token.Valid {
		return nil, errors.New("error")
	}
	return claims, nil
}

func RefreshCheck(refresh string) (string, error) {
	refreshTk, err := decrypt(refresh)
	if err != nil {
		return "", err
	}

	tk, err := ParseToken(refreshTk)
	if err != nil {
		return "", err
	}

	return tk.Id, nil
}

func validateSignedMethod(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}
	return JwtSecretKey, nil
}

//============================== encrypt text zone ==============================//

// Encrypt method is to encrypt or hide any classified text
func encrypt(text string) string {
	enc := base64.RawStdEncoding.EncodeToString([]byte(text))

	return enc
}

// Decrypt method is to extract back the encrypted text
func decrypt(text string) (string, error) {
	dec, err := base64.RawStdEncoding.DecodeString(text)
	if err != nil {
		return "", err
	}

	return string(dec), nil
}

func encodeBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func decodeBase64(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}
