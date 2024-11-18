package env

import "os"

func Env(key string) string {
	return os.Getenv(key)
}
