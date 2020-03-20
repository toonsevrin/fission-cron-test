package main

// Whenever the handler is called, a POST on the url specified in the configmap is called.

import (
	"fmt"
	"io"
	"net/http"
	"io/ioutil"
)

func get_url() (string, error) {
	url, err := ioutil.ReadFile("/configs/default/fission-cron-test/url")
	return string(url), err
}
// handle is the actual processor of the function
func handle(w http.ResponseWriter, r *http.Request) error {
	url, err := get_url()
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "", nil)
	if err != nil {
		return err
	}
	io.Copy(w, resp.Body)
	return nil
}

// Handler is the entry point for this fission function
func Handler(w http.ResponseWriter, r *http.Request) {
        err := handle(w, r)
        if err != nil {
                  w.WriteHeader(http.StatusInternalServerError)
                  w.Write([]byte(fmt.Sprintf("500 - Something bad happened!\nError: %s", err.Error())))
                  return
        }
}

