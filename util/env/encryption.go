package env

func GetEncryptionKey() string {
	key := getEnvVar("ENCRYPTION_KEY")
	if key == "" {
		return "defaultkey"
	}
	return key
}
