package app

import "fmt"

func MakeBold(message string) string {
	return fmt.Sprintf("**%s**", message)
}
