package main

import (
	"log"
	"net/http"
	"github.com/bobsar0/autotrade/app"
)

//Program starts here
func main() {

	s := app.NewSession() //Initializes new session
	h := app.NewAppHandler(s) //Passes the session to initialize a new instance of appHandler
	
	//Starting the goroutines
	go app.GetTicker(s.GetTickerChan)
	go app.GetBalance(s.GetBalanceChan)
	go app.PlaceOrder(s.PlaceOrderChan)

	log.Println("Connecting to server on port 8000...")
	log.Fatalln(http.ListenAndServe(":8000", h)) //Set listening port (:8080). Handler is h indicating that chi router is used. log.Fatal checks for error and outputs if any.
}
