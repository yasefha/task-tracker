package cli

const (
	colorReset = "\033[0m"
	colorBold  = "\033[1m"

	colorRed     = "\033[31m"
	colorGreen   = "\033[32m"
	colorYellow  = "\033[33m"
	colorBlue    = "\033[34m"
	colorMagenta = "\033[35m"
	colorCyan    = "\033[36m"
	colorGray    = "\033[90m"
)

func bold(text string) string {
	return colorBold + text + colorReset
}

func green(text string) string {
	return colorGreen + text + colorReset
}

func red(text string) string {
	return colorRed + text + colorReset
}

func yellow(text string) string {
	return colorYellow + text + colorReset
}

func cyan(text string) string {
	return colorCyan + text + colorReset
}

func magenta(text string) string {
	return colorMagenta + text + colorReset
}

func blue(text string) string {
	return colorBlue + text + colorReset
}

func gray(text string) string {
	return colorGray + text + colorReset
}
