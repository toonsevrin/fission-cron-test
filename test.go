package main

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
)
func main() {
		req := httptest.NewRequest("GET","/", nil)
		w := httptest.NewRecorder()

		Handler(w, req)
		body, _ := ioutil.ReadAll(w.Body)
		fmt.Printf("Body: %s", string(body))
}