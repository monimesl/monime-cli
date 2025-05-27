package browser

func open(url string) (Command, error) {
	programs := []string{"xdg-open", "x-www-browser", "www-browser"}
	for _, provider := range programs {
		if _, err := exec.LookPath(provider); err == nil {
			return runCmd(provider, url)
		}
	}
	return &exec.Error{Name: strings.Join(programs, ","), Err: exec.ErrNotFound}
}
