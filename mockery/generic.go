package mockery

import (
	"math/rand"
	"os"
	"time"
	"unicode"
)

// Config read from environment variables
type Config map[string]string

// GetConfigFromEnv gets config from environment
func GetConfigFromEnv() Config {	
	return Config{
		"apiKey": os.Getenv("apiKey"),
	}
}

// ConvertToMockery conVeRTs TexT tO MOcKerY
func ConvertToMockery(str string) string {
	var convertedStr []rune
	for i := 0; i < len(str); i++ {
		rand.Seed(time.Now().UnixNano())
		if rand.Intn(2) != 0 {
			convertedStr = append(convertedStr, unicode.ToUpper(rune(str[i])))
		} else {
			convertedStr = append(convertedStr, unicode.ToLower(rune(str[i])))
		}
	}
	return string(convertedStr)
}