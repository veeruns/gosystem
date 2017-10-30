package gosystem

import (
	"net/http"
)

func StartServer(addr string) error {
	http.HandleFunc("/", rootHandler)
	err := http.ListenAndServe(addr, nil)
	return err
}

func rootHandler(w http.ResponseWriter, r *http.Request) {

}
