package login

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/monime-lab/gwater"
	"github.com/monimesl/monime-cli/internal/browser"
	"github.com/monimesl/monime-cli/internal/prompts"
	"github.com/monimesl/monime-cli/internal/text"
	"golang.org/x/crypto/nacl/box"
)

const (
	callbackPort   = "37891"
	callbackPath   = "/callback"
	remoteLoginURL = "https://cli.monime.io/login"
	callbackURL    = "http://localhost:" + callbackPort + callbackPath
)

type Flow struct {
	verifier  tokenVerifier
	decrypter tokenDecrypter
}

func (f *Flow) Run(ctx context.Context) (Token, error) {
	publicKey, privateKey, err := box.GenerateKey(rand.Reader)
	if err != nil {
		return Token{}, err
	}
	state := gwater.RandomURLSafeString(64)
	encodedPublicKey := base64.StdEncoding.EncodeToString(publicKey[:])
	securedToken, err := f.launchInteraction(ctx, encodedPublicKey, state)
	if err != nil {
		return Token{}, err
	}
	fmt.Println("🔏 Decrypting token...")
	tokenBytes, err := f.decrypter.decrypt(securedToken, privateKey)
	if err != nil {
		return Token{}, err
	}
	securedToken.bytes = tokenBytes
	fmt.Println("🔏 Verifying token signature...")
	if err = f.verifier.verify(securedToken); err != nil {
		return Token{}, err
	}
	token, err := securedToken.extractToken()
	if err != nil {
		return Token{}, err
	}
	if token.State != state {
		return Token{}, errors.New("token invalid state from callback")
	}
	return token, nil
}

//nolint:cyclop
func (f *Flow) launchInteraction(ctx context.Context, publicKey, state string) (SecuredToken, error) {
	tokenCh := make(chan SecuredToken)
	srv := &http.Server{Addr: ":" + callbackPort}
	http.HandleFunc(callbackPath, func(w http.ResponseWriter, r *http.Request) {
		setCORSHeaders(w, r)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}
		secret := r.FormValue("secret")
		signature := r.FormValue("signature")
		if secret == "" || signature == "" {
			http.Error(w, "Missing secret/signature", http.StatusBadRequest)
			return
		}
		tokenCh <- SecuredToken{Secret: secret, Signature: signature}
	})
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Unable to start login flow server: %v", err)
		}
	}()
	defer func() {
		shutdownCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
		defer cancel()
		_ = srv.Shutdown(shutdownCtx)
	}()
	authURL, err := buildAuthURL(publicKey, state)
	if err != nil {
		return SecuredToken{}, err
	}
	text.PrintStart(fmt.Sprintf("%s to open the URL in your browser to initiate the login.",
		text.Format("Press Enter", text.FormatOptions{Color: "white", Bold: true})))
	if ok, _ := prompts.EnterKey(ctx); !ok {
		os.Exit(0)
	}
	fmt.Println("🔐 Opening browser for login...")
	if _, err = browser.OpenURL(ctx, authURL); err != nil {
		return SecuredToken{}, fmt.Errorf("failed to open browser: %w", err)
	}
	fmt.Printf("🌐 Listening on %s\n", callbackURL)
	fmt.Println("🕒 Waiting for login to complete...")
	select {
	case token := <-tokenCh:
		fmt.Println("✅ Token received.")
		return token, nil
	case <-ctx.Done():
		return SecuredToken{}, ctx.Err()
	case <-time.After(2 * time.Minute):
		return SecuredToken{}, fmt.Errorf("timeout waiting for login callback")
	}
}

func setCORSHeaders(w http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")
	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Private-Network", "true")
}

func buildAuthURL(publicKey, state string) (string, error) {
	authURL, err := url.Parse(remoteLoginURL)
	if err != nil {
		return "", err
	}
	query := authURL.Query()
	query.Set("state", state)
	query.Set("publicKey", publicKey)
	query.Set("callbackURL", callbackURL)
	authURL.RawQuery = query.Encode()
	return authURL.String(), nil
}
