package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"regexp"
)

func failHEAD(w http.ResponseWriter, r *http.Request) {
	if r.Method == "HEAD" {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	redirect(w, r)
}

func flaky(w http.ResponseWriter, r *http.Request) {
	random, err := rand.Int(rand.Reader, big.NewInt(20))
	if err != nil {
		panic(err)
	}
	if random.Cmp(big.NewInt(0)) == 0 {
		// 1-in-20 odds of success
		redirect(w, r)
		return
	}
	w.WriteHeader(http.StatusServiceUnavailable)
}

func redirect(w http.ResponseWriter, r *http.Request) {
	re := regexp.MustCompile(`[^/]*$`)
	requestedFile := re.Find([]byte(r.URL.String()))

	redirectURL := fmt.Sprintf("http://cdi-file-host.cdi/%s", requestedFile)
	http.Redirect(w, r, redirectURL, 301)
}

func main() {
	http.HandleFunc("/forbidden-HEAD/", failHEAD)
	http.HandleFunc("/flaky/", flaky)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
