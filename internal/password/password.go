package password

import (
	"math/rand"
	"strings"
	"time"
)

var (
	lowerCharSet   = "abcdedfghijklmnopqrst"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!@#$%&*"
	numberSet      = "0123456789"
)

func Generate(passwordLength int, includeUpperCase, includeNum, includeSymbol bool) (string, error) {
	rand.Seed(time.Now().Unix())
	divider := 1
	var lowerCaseDistribution int
	var upperCaseDistribution int
	var numDistribution int
	var symbolDistribution int
	if includeUpperCase {
		divider += 1
	}
	if includeNum {
		divider += 1
	}
	if includeSymbol {
		divider += 1
	}

	var password strings.Builder

	res := passwordLength / divider
	remainder := passwordLength % divider

	lowerCaseDistribution = res + remainder

	for i := 0; i < lowerCaseDistribution; i++ {
		random := rand.Intn(len(lowerCharSet))
		password.WriteString(string(lowerCharSet[random]))
	}
	if includeUpperCase {
		upperCaseDistribution = passwordLength / divider
		for i := 0; i < upperCaseDistribution; i++ {
			random := rand.Intn(len(upperCharSet))
			password.WriteString(string(upperCharSet[random]))
		}
	}

	if includeNum {
		numDistribution = passwordLength / divider
		for i := 0; i < numDistribution; i++ {
			random := rand.Intn(len(numberSet))
			password.WriteString(string(numberSet[random]))
		}
	}

	if includeSymbol {
		symbolDistribution = passwordLength / divider
		for i := 0; i < symbolDistribution; i++ {
			random := rand.Intn(len(specialCharSet))
			password.WriteString(string(specialCharSet[random]))
		}
	}

	pswd := []rune(password.String())
	rand.Shuffle(len(pswd), func(i, j int) {
		pswd[i], pswd[j] = pswd[j], pswd[i]
	})

	return string(pswd), nil
}
