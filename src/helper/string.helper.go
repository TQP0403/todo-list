package helper

import "os"

func GetDefaultString(val, defaultVal string) string {
	if len(val) == 0 {
		return defaultVal
	}
	return val
}

func GetDefaultEnv(key, defaultVal string) string {
	val := os.Getenv(key)
	if len(val) == 0 {
		return defaultVal
	}
	return val
}
