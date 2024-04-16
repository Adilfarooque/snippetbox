package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	/*Define a new command-line flag with the name "addr" a default value of ":4000"
	and some short help text explaining what the flag controls. The value of the flag
	will be stored in the addr varable at runtime
	*/
	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()
	//Create a file server which serves files out of the "./ui/static" directory
	//Note that the path given to the http.Dir function is relative to the project
	//directory root.
	fileServer := http.FileServer(http.Dir("C:/Users/USER/Desktop/code/snippetbox/ui/static"))
	//User the mux.Handle() fuction to register the file server as the handler
	//all URL paths that start with "/static/" . For matching paths, we strip the
	//"/static" prefix before the request reaches the server
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Printf("Starting server on %s",*addr)
	if err := http.ListenAndServe(*addr, mux); err != nil {
		log.Fatal(err)
	}
}
