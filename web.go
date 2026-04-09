package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/{$}", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "html/index.html")
	})
	router.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "images/dontwin.svg")
	})
	router.Handle("/html/", http.StripPrefix("/html", http.FileServer(http.Dir("html"))))
	router.Handle("/images/", http.StripPrefix("/images", http.FileServer(http.Dir("images"))))
	router.Handle("/", notFound(router))

	server := http.Server{Addr: ":8080", Handler: logger(router), ReadHeaderTimeout: time.Second}
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}

func notFound(mux *http.ServeMux) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var allowed []string
		method := r.Method
		_, current := mux.Handler(r)
		for _, method := range methods() {
			r.Method = method
			if _, pattern := mux.Handler(r); pattern != current {
				allowed = append(allowed, method)
			}
		}
		r.Method = method
		if len(allowed) != 0 {
			w.Header().Set("Allow", strings.Join(allowed, ", "))
			w.WriteHeader(http.StatusMethodNotAllowed)
			io.WriteString(w, "this is not the method you are looking for...")
			return
		}
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "<div style=\"margin:auto; text-align:center\">")
		io.WriteString(w, "<p>this is not the page you are loooking for...</p>")
		io.WriteString(w, "<p>go about your business</p><p>move along</p></div>")
	})
}

func methods() []string {
	return []string{
		http.MethodGet,
		http.MethodHead,
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
		http.MethodConnect,
		http.MethodOptions,
		http.MethodTrace,
	}
}

type statusRecorder struct {
	http.ResponseWriter

	status int
}

// WriteHeader overrides std WriteHeader to save response code.
func (rec *statusRecorder) WriteHeader(code int) {
	rec.status = code
	rec.ResponseWriter.WriteHeader(code)
}

// Logger is a logging middleware that logs useragent, RemoteAddr, Method, Host, Path and response.Status to stdlib log.
func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		rec := statusRecorder{w, http.StatusOK}
		next.ServeHTTP(&rec, r)
		// remote := strings.Split(r.RemoteAddr, ":")[0]
		remote := r.RemoteAddr
		if r.Header.Get("X-Forwarded-For") != "" {
			remote = r.Header.Get("X-Forwarded-For")
		}
		details := fmt.Sprintf(
			"%s %s%s %d %s %s %s",
			r.Method,
			r.Host,
			r.URL.Path,
			rec.status,
			remote,
			time.Since(now).String(),
			r.UserAgent(),
		)
		log.Println(details)
	})
}
