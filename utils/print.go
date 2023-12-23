package utils

import (
	"fmt"
)

func PrintDefault(message string, defaultVal any) string {
	return fmt.Sprintf("%s [%v]", message, defaultVal)
}
