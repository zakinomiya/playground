package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"golang.org/x/sync/errgroup"
)

func main() {
	eg := errgroup.Group{}
	eg.Go(runServer)
	eg.Go(runProxy)
	fmt.Println("Server and Proxy are running")
	if err := eg.Wait(); err != nil {
		fmt.Println(err)
	}
}

func runServer() error {
	mux := http.NewServeMux()

	mux.Handle("/hello", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	}))
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, _ := httputil.DumpRequest(r, true)
		fmt.Println(string(b))
		w.Write([]byte("Hello!"))
	}))

	return http.ListenAndServe(":8080", mux)
}

func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}

func joinURLPath(a, b *url.URL) (path, rawpath string) {
	if a.RawPath == "" && b.RawPath == "" {
		return singleJoiningSlash(a.Path, b.Path), ""
	}
	// Same as singleJoiningSlash, but uses EscapedPath to determine
	// whether a slash should be added
	apath := a.EscapedPath()
	bpath := b.EscapedPath()

	aslash := strings.HasSuffix(apath, "/")
	bslash := strings.HasPrefix(bpath, "/")

	switch {
	case aslash && bslash:
		return a.Path + b.Path[1:], apath + bpath[1:]
	case !aslash && !bslash:
		return a.Path + "/" + b.Path, apath + "/" + bpath
	}
	return a.Path + b.Path, apath + bpath
}

func runProxy() error {
	mux := http.NewServeMux()
	s, err := url.Parse("http://localhost:8080")
	if err != nil {
		panic(err)
	}
	rp := httputil.NewSingleHostReverseProxy(s)
	rp.Director = func(req *http.Request) {
		target := req.URL
		targetQuery := target.RawQuery
		req.URL.Scheme = target.Scheme
		if target.Scheme == "" {
			req.URL.Scheme = "http"
		}
		req.URL.Host = "localhost:8080"
		p, rp := joinURLPath(s, target)
		fmt.Println(p, rp)
		req.URL.Path, req.URL.RawPath = strings.TrimPrefix(p, "/picking"), rp
		if targetQuery == "" || req.URL.RawQuery == "" {
			req.URL.RawQuery = targetQuery + req.URL.RawQuery
		} else {
			req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
		}
	}

	mux.Handle("/hello", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Proxy!"))
	}))
	mux.Handle("/picking/", rp)

	return http.ListenAndServe(":8081", mux)
}
