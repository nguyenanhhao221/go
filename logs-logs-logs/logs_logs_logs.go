package logs

var tableLookUp = map[rune]string{}

// Application identifies the application emitting the given log.
func Application(log string) string {
	// panic("Please implement the Application() function")
	tableLookUp['❗'] = "recommendation"
	tableLookUp['🔍'] = "search"
	tableLookUp['☀'] = "weather"
	for _, char := range log {
		value, exist := tableLookUp[char]
		if exist {
			return value
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	panic("Please implement the Replace() function")
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	panic("Please implement the WithinLimit() function")
}
