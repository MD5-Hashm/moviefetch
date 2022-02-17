package chars

var (
	chars = map[string]string{
		"up":      "↑",
		"down":    "↓",
		"warning": "(Low)"}
)

func Get(char string) string {
	return chars[char]
}
