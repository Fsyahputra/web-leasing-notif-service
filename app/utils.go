package app

import "fmt"

func makeBold(message string) string {
	return fmt.Sprintf("**%s**", message)
}
