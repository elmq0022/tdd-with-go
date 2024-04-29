package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run(
		"compare server speeds, return url of the fastest one",
		func(t *testing.T) {
			slowServer := makeDelayedServer(20 * time.Millisecond)
			fastServer := makeDelayedServer(0 * time.Millisecond)

			defer slowServer.Close()
			defer fastServer.Close()

			slowURL := slowServer.URL
			fastURL := fastServer.URL

			want := fastURL
			got, _ := Racer(slowURL, fastURL)

			if got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		},
	)

	t.Run("return an error if a server doesn't respond in 10s", func(t *testing.T) {
		serverA := makeDelayedServer(20 * time.Millisecond)
		serverB := makeDelayedServer(30 * time.Millisecond)

		defer serverA.Close()
		defer serverB.Close()

		_, err := ConfigurableRacer(serverA.URL, serverB.URL, 10*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})

}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
	return server
}
