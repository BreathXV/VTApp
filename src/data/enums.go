package data

type Mode int8

// Modes at which the application will notify
// the user of the results of the file report
// contents.
const (
	MODE_INFO      Mode = 0
	MODE_DANGER    Mode = 1
	MODE_MALICIOUS Mode = 2
)
