package monimeapis

import (
	"sync"

	webclient "github.com/monime-lab/web-client-go"
)

const (
	baseUrl = "https://cli.monime.io/api"
)

var (
	webClientOnce sync.Once
	webClient     webclient.Client
)

func getWebClient() webclient.Client {
	webClientOnce.Do(func() {
		cl := webclient.New(nil,
			webclient.WithBaseUrl(baseUrl),
			webclient.EnableTracing(),
			webclient.EnableDebug(),
		)
		webClient = cl
	})
	return webClient
}
