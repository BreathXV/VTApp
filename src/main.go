package main

// On application start, constantly checks specified folder for new files/folders,
// then parses them to check file size

type Application struct {
	name    string
	author  string
	license string
	version float32
}

func main() {
	VTApp := Application{
		name:    "VTApp",
		author:  "BreathXV",
		license: "N/a",
		version: 0.1,
	}
}
