package token_handler

import (
	error_handler2 "api-authenticator-proxy/util/error_handler"
	tokenError "api-authenticator-proxy/util/error_handler/token"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io"
)

func Generate(id string, passphrase string) (string, error_handler2.StatusError) {
	data := map[string]string{
		"id":         id,
		"passphrase": passphrase,
	}

	jsonData, err1 := json.Marshal(data)
	if err1 != nil {
		return "", error_handler2.UnexpectedError("unexpected error : Error in provided data")
	}

	gcm, err := createGCMBloc()
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err1 = io.ReadFull(rand.Reader, nonce); err1 != nil {
		return "", error_handler2.UnexpectedError("unexpected error : Error in nonce generation")
	}

	ciphertext := gcm.Seal(nonce, nonce, jsonData, nil)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func Decrypt(token string) (string, error_handler2.StatusError) {
	decodedToken, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return "", tokenError.InvalidTokenError()
	}
	gcm, err1 := createGCMBloc()
	if err1 != nil {
		return "", err1
	}

	nonceSize := gcm.NonceSize()
	if len(decodedToken) < nonceSize {
		return "", tokenError.InvalidTokenError()
	}

	nonce, ciphertext := decodedToken[:nonceSize], decodedToken[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", tokenError.InvalidTokenError()
	}

	var decryptedStruct map[string]string
	err = json.Unmarshal(plaintext, &decryptedStruct)
	if err != nil {
		return "", tokenError.InvalidTokenError()
	}
	return decryptedStruct["id"], nil
}

func createGCMBloc() (cipher.AEAD, error_handler2.StatusError) {
	var key = []byte("example key 1234")

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, error_handler2.UnexpectedError("Error in key generation")
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, error_handler2.UnexpectedError("Error in GCM generation")
	}
	return gcm, nil
}
