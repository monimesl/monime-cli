package main

import (
	"github.com/monimesl/monime-cli/cmds"
	"github.com/monimesl/monime-cli/pkg/browser"
	"net/http"
	"time"
)

func main() {
	cmds.ExecuteRootCmd()
	//keyringTest()
}

func keyringTest() {
	http.HandleFunc("/callback2", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <title>Submitting...</title>
</head>
<body>
  <form id="authForm" action="/callback3" method="POST">
    <!-- Optional hidden input -->
    <input type="hidden" name="token" value="your_encrypted_token_here">
  </form>

  <script>
    document.getElementById("authForm").submit();
  </script>

  <p>Submitting authentication... If this page doesn't close, you may do so manually.</p>
</body>
</html>
`))
	})
	http.HandleFunc("/callback3", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
      <html><body>
      <p>This page will close in 2 seconds.</p>
      <script>setTimeout(() => window.close(), 2000);</script>
      </body></html>
    `))
	})
	go http.ListenAndServe(":37891", nil)
	browser.OpenURL("http://localhost:37891/callback2")
	time.Sleep(10 * time.Second)
}
