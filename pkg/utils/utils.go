package utils

func Must(e any) {
	if e != nil {
		panic(e)
	}
}
