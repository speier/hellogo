package learn

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func WebServer() string {
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello from Go webserver!")
		})
		log.Fatal(http.ListenAndServe(":3000", nil))
	}()

	res, err := http.Get("http://localhost:3000")
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	return string(body)
}
