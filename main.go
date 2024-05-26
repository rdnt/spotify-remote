package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	authn "github.com/rdnt/rdnt/pkg/oauth"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2"
	spotifyOauth "golang.org/x/oauth2/spotify"

	"spotify-seeker/keyboard"
)

func main() {
	host := os.Getenv("SERVER_HOST")
	strPort := os.Getenv("SERVER_PORT")

	if host == "" {
		host = "localhost"
	}

	if strPort == "" {
		strPort = "8080"
	}

	port, err := strconv.Atoi(strPort)
	if err != nil {
		log.Fatal(err)
	}

	spotifyClientId := os.Getenv("SPOTIFY_CLIENT_ID")
	spotifyClientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")
	spotifyRedirectUrl := os.Getenv("SPOTIFY_REDIRECT_URL")

	spotifyConf := &oauth2.Config{
		ClientID:     spotifyClientId,
		ClientSecret: spotifyClientSecret,
		Scopes: []string{
			spotifyauth.ScopeUserReadPrivate,
			spotifyauth.ScopeUserReadPlaybackState,
			spotifyauth.ScopeUserModifyPlaybackState,
		},
		Endpoint:    spotifyOauth.Endpoint,
		RedirectURL: spotifyRedirectUrl,
	}

	spotifyTokenProv := authn.NewInMemoryTokenProvider()

	spotifyAuthn, err := authn.NewAuthn("Spotify", spotifyConf, spotifyTokenProv)
	if err != nil {
		log.Fatal(err)
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	err = r.SetTrustedProxies(nil)
	if err != nil {
		log.Fatal(err)
	}

	r.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	oauth := r.Group("/oauth")
	oauth.Use(gin.Logger())

	oauth.GET("/spotify/callback", func(c *gin.Context) {
		err := spotifyAuthn.ExtractToken(c.Request)
		if errors.Is(err, authn.ErrExchanged) {
			return
		} else if err != nil {
			log.Println("spotify", err)
			return
		}

		c.String(http.StatusOK, "Spotify successfully authenticated. You may close this window.")
		log.Print("Spotify client authenticated successfully.")
	})

	go func() {
		err := r.Run(fmt.Sprintf("%s:%d", host, port))
		if !errors.Is(err, http.ErrServerClosed) && err != nil {
			log.Fatal(err)
		}
	}()

	spotifyHttpClient, err := spotifyAuthn.Client()
	if err != nil {
		log.Fatal(err)
	}

	spotifyClient := spotify.New(
		spotifyHttpClient,
		spotify.WithRetry(true),
	)
	kb := keyboard.New(func(code int) {
		switch code {
		case keyboard.NextTrack:
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			err := spotifyClient.Next(ctx)
			if err != nil {
				fmt.Println("Next failed:", err)
				return
			}

			fmt.Println("Next track")
		case keyboard.PlayPause:
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			curr, err := spotifyClient.PlayerCurrentlyPlaying(ctx)
			if err != nil {
				fmt.Println("CurrentlyPlaying failed:", err)
				return
			}

			if !curr.Playing {
				err := spotifyClient.Play(ctx)
				if err != nil {
					fmt.Println("Play failed:", err)
					return
				}

				fmt.Println("Resumed")
			} else {
				err := spotifyClient.Pause(ctx)
				if err != nil {
					fmt.Println("Pause failed:", err)
					return
				}

				fmt.Println("Paused")
			}
		}
	}, nil)

	kb.Capture()

	done := make(chan os.Signal, 10)
	signal.Notify(done, os.Interrupt, os.Kill, syscall.SIGTERM)

	<-done

	kb.Stop()
}
