package utils

import (
	"bytes"
	"os"
	"unicode"
)

func GetEnvVar(vn string) string {
    var varExtension string
    if os.Getenv("PRODUCTION") == "1" {
        varExtension = "_PRODUCTIOn" 
    }

    return os.Getenv(vn + varExtension)
}

func ConvertCamelToSnake(camelCase string) string {
	var buffer bytes.Buffer

	for i, char := range camelCase {
		if unicode.IsUpper(char) {
			if i > 0 {
				buffer.WriteRune('_')
			}
			buffer.WriteRune(unicode.ToLower(char))
		} else {
			buffer.WriteRune(char)
		}
	}

	return buffer.String()
}

// Enmurates over a map
func Enumerate(m map[string]interface{}) []Pair {
	pairs := make([]Pair, len(m))
	i := 0
	for k, v := range m {
		pairs[i] = Pair{k, v}
		i++
	}
	return pairs
}

type Pair struct {
	Key   string
	Value interface{}
}
