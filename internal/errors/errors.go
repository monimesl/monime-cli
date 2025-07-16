package errors

import (
	"errors"
	"fmt"
)

var (
	ErrCliSilent       = errors.New("cli silent")
	ErrNoActivateSpace = errors.New("no activate space")
)

func PrintLoginHint() {
	fmt.Println("\033[1;31m❌  Authentication Required\033[0m")
	fmt.Println("You need to be logged in to use this command.")

	fmt.Println()
	fmt.Println("👉 To log in, run:")
	fmt.Println("   \033[1;36mmonime account login\033[0m")

	fmt.Println()
	fmt.Println("🛟 Need help? Visit: \033[4;34mhttps://docs.monime.io/cli\033[0m")
	fmt.Println()
}
