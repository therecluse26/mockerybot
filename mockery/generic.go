package mockery

import (
	"math/rand"
	"os"
	"strings"
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

var ApologiesWithName = []string {
	"I'm sorry, {name}",
	"Please forgive me, {name}",
	"I prostrate myself before you and beg your forgiveness, {name}",
	"{name}, I really messed up",
	"I'm so sorry {name}, please don't hurt me",
}

var ApologiesNoName = []string {
	"Sorrrrryyy",
	"I messed up, sorry about that",
	"That wasn't cool, sorry...",
	"*Incoherent Samurai Screaming*",
	"It's all my fault",
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

func MakeApology(name string) string {

	var apology string

	if len(strings.TrimSpace(name)) > 0 {
		s := rand.NewSource(time.Now().Unix())
		r := rand.New(s)
		idx := r.Intn(len(ApologiesWithName))
		apology = ApologiesWithName[idx]
		apology = strings.Replace(apology, "{name}", name, -1)

	} else {
		s := rand.NewSource(time.Now().Unix())
		r := rand.New(s)
		idx := r.Intn(len(ApologiesNoName))
		apology = ApologiesNoName[idx]
	}
	return apology
}