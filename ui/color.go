package ui

import "fmt"

const (
	Reset  = "\033[0m"
	Bold   = "\033[1m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Cyan   = "\033[36m"
)

func GreenText(s string) string {
	return Green + s + Reset
}

func RedText(s string) string {
	return Red + s + Reset
}

func YellowText(s string) string {
	return Yellow + s + Reset
}

func CyanText(s string) string {
	return Cyan + s + Reset
}

func BoldText(s string) string {
	return Bold + s + Reset
}

func PrintfColor(color string, format string, a ...interface{}) {
	fmt.Print(color)
	fmt.Printf(format, a...)
	fmt.Print(Reset)
}
