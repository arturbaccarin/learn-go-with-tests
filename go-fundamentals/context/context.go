package context

import (
	"fmt"
	"net/http"
)

type Store interface {
	Fetch() string
	Cancel()
}

// we'll use the package context to help us manage long-running processes.
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		data := make(chan string, 1)

		go func() {
			data <- store.Fetch()
		}()

		select {
		case d := <-data:
			fmt.Fprint(w, d)
		/*
			context has a method Done() which returns a channel which gets sent a signal
			when the context is "done" or "cancelled". We want to listen to that signal and
			call store.Cancel if we get it but we want to ignore it if our Store manages to
			Fetch before it.
		*/
		case <-ctx.Done():
			store.Cancel()
		}
	}
}
