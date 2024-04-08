package token_handler

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
)

func Generate(id string, passphrase string) (string, error) {
	data := map[string]string{
		"id":         id,
		"passphrase": passphrase,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("error in the provided data")
	}

	gcm, err := createGCMBloc()
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", fmt.Errorf("unexpected error : Please try again")
	}

	ciphertext := gcm.Seal(nonce, nonce, jsonData, nil)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func Decrypt(token string) (string, error) {
	decodedToken, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return "", fmt.Errorf("error decoding the token: invalid token provided")
	}
	gcm, err := createGCMBloc()
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(decodedToken) < nonceSize {
		return "", fmt.Errorf("error decoding the token: invalid token provided")
	}

	nonce, ciphertext := decodedToken[:nonceSize], decodedToken[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", fmt.Errorf("error decoding the token: invalid token provided")
	}

	var decryptedStruct map[string]string
	err = json.Unmarshal(plaintext, &decryptedStruct)
	if err != nil {
		return "", fmt.Errorf("error extracting the id from the token")
	}
	return decryptedStruct["id"], nil
}

func createGCMBloc() (cipher.AEAD, error) {
	var key = []byte("example key 1234")

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("invalid key : Please check your key")
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("invalid key : Please check your key")
	}
	return gcm, nil
}
