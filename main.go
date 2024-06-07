package main

import (
	"log/slog"
	"net/http"
)

func main() {
	slog.Info("starting server on 3000")
	err := http.ListenAndServe(":3000", hndlr())
	if err != nil && err != http.ErrServerClosed {
		slog.Error("we got an error", "err", err)
	}
}

func hndlr() http.Handler {
	hndl := http.NewServeMux()
	hndl.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Referer() == "/redirect" {
			w.Write([]byte("redirected from /redirect"))
			return
		}
		w.Write([]byte("hello world"))
	})

	hndl.HandleFunc("/redirect", func(w http.ResponseWriter, r *http.Request) {
		referer := r.Referer()
		slog.Info("referer", "referer", referer)
		slog.Info("redirecting request to empty string")
		if r.Referer() == "/redirect" {
			w.Write([]byte("just kidding, we just came from there -- infinite loop"))
			return
		}
		http.Redirect(w, r, "", http.StatusFound)
	})

	hndl.HandleFunc("/red-found", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("redirecting request to /redirect")
		http.Redirect(w, r, "/redirect", http.StatusFound)
	})

	hndl.HandleFunc("/red-perm", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("redirecting request to /redirect")
		http.Redirect(w, r, "/redirect", http.StatusPermanentRedirect)
	})

	hndl.HandleFunc("/red-mvd-perm", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("redirecting request to /redirect")
		http.Redirect(w, r, "/redirect", http.StatusMovedPermanently)
	})

	return hndl
}
