package token_handler

import (
	"api-authenticator-proxy/util/env"
	"api-authenticator-proxy/util/error_handler"
	tokenError "api-authenticator-proxy/util/error_handler/token"
	"api-authenticator-proxy/util/log"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
)

func Generate(id string, passphrase string) (string, error_handler.StatusError) {
	data := map[string]string{
		"id":         id,
		"passphrase": passphrase,
	}
	log.Debug("generating token from : ", data)
	jsonData, err1 := json.Marshal(data)
	if err1 != nil {
		log.Error(fmt.Errorf("unexpected error : Error in provided data : %v", err1))
		return "", error_handler.UnexpectedError("unexpected error : Error in provided data")
	}

	gcm, err := createGCMBloc()
	if err != nil {
		log.Error(fmt.Errorf("unexpected error : Error in GCM generation : %v", err))
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err1 = io.ReadFull(rand.Reader, nonce); err1 != nil {
		log.Error(fmt.Errorf("unexpected error : Error in nonce generation : %v", err1))
		return "", error_handler.UnexpectedError("unexpected error : Error in nonce generation")
	}

	ciphertext := gcm.Seal(nonce, nonce, jsonData, nil)

	generatedToken := base64.StdEncoding.EncodeToString(ciphertext)

	log.Debug("generated token : ", generatedToken)
	return generatedToken, nil
}

func Decrypt(token string) (string, error_handler.StatusError) {
	decodedToken, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		log.Error(fmt.Errorf("unexpected error : Error in decoding token : %v", err))
		return "", tokenError.InvalidTokenError()
	}
	gcm, err1 := createGCMBloc()
	if err1 != nil {
		log.Error(fmt.Errorf("unexpected error : Error in GCM generation : %v", err1))
		return "", err1
	}

	nonceSize := gcm.NonceSize()
	if len(decodedToken) < nonceSize {
		log.Error(fmt.Errorf("unexpected error in token : %v", tokenError.InvalidTokenError()))
		return "", tokenError.InvalidTokenError()
	}

	nonce, ciphertext := decodedToken[:nonceSize], decodedToken[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		log.Error(fmt.Errorf("unexpected error in token : %v", tokenError.InvalidTokenError()))
		return "", tokenError.InvalidTokenError()
	}

	var decryptedStruct map[string]string
	err = json.Unmarshal(plaintext, &decryptedStruct)
	if err != nil {
		log.Error(fmt.Errorf("unexpected error in token : %v", tokenError.InvalidTokenError()))
		return "", tokenError.InvalidTokenError()
	}

	log.Debug("decrypted token : ", decryptedStruct["id"])
	return decryptedStruct["id"], nil
}

func createGCMBloc() (cipher.AEAD, error_handler.StatusError) {
	encryptionKey := env.GetEncryptionKey()
	var key = []byte(encryptionKey)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, error_handler.UnexpectedError("Error in key generation")
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, error_handler.UnexpectedError("Error in GCM generation")
	}
	return gcm, nil
}
