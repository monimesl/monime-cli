package browser

func open(url string) (Command, error) {
	return runCmd("open", url)
}
