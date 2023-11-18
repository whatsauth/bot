package main

import (
	"net/http"

	"github.com/whatsauth/bot"
)

func main() {
	http.HandleFunc("/", bot.PostWithDB)
	http.ListenAndServe(":8080", nil)
}
