package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/availability/room/book", book)
	http.HandleFunc("/availability/room/4.1", bookRoom)
	http.HandleFunc("/availability/room/4.2", bookRoom)
	http.HandleFunc("/availability", availability)
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil)
}

func homePage(responseWritter http.ResponseWriter, request *http.Request) {
	responseWritter.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(responseWritter, "Welcome to our booking system!<br><a href=\"/availability\">List all available rooms</a>")
}

func availability(responseWritter http.ResponseWriter, request *http.Request) {
	responseWritter.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(responseWritter, `
		<ol>
			<a href="/availability/room/4.1">
				<li>
					Room 4.1
				</li>
			</a>
			<a href="/availability/room/4.2">
				<li>
					Room 4.2
				</li>
			</a>
		</ol>
		`)
}

func bookRoom(responseWritter http.ResponseWriter, request *http.Request) {
	responseWritter.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(responseWritter, `
		<h1>Would you like to confirm your booking</h1>
		<form action="./book" method="post">
			<button type="submit">
				Confirm
			</button>
		</form>
		`)
}

func book(responseWritter http.ResponseWriter, request *http.Request) {
	responseWritter.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(responseWritter, `<h1>Thank you for your booking</h1>`)
}
