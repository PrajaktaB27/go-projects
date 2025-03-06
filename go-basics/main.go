package main

import (
	"net/http"
)

// consuming the net -> http package
// type server struct {
// 	addr string
// }

// func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	// w.Write([]byte("Hi"))
// 	switch r.Method {
// 	case http.MethodGet:
// 		switch r.URL.Path {
// 		case "/":
// 			w.Write([]byte("index page"))
// 			return
// 		case "/users":
// 			w.Write([]byte("users page"))
// 			return
// 		}
// 	default:
// 		w.Write([]byte("404 page not found"))
// 		return
// 	}
// }

// type api struct {
// 	addr string
// }

// func (s *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	// w.Write([]byte("Hi"))
// 	switch r.Method {
// 	case http.MethodGet:
// 		switch r.URL.Path {
// 		case "/":
// 			w.Write([]byte("GET method"))
// 			return
// 		case "/index":
// 			w.Write([]byte("GET index"))
// 			return
// 		}
// 	case http.MethodPost:
// 		w.Write([]byte("GET method"))
// 		// default:
// 		// 	w.Write([]byte("404 page not found"))
// 		// 	return
// 	}
// }

// func (a *api) getUsersHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("users list - "))
// }

// func (a *api) createUsersHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("created user - "))
// }

func main() {
	// s := &server{addr: ":8080"}
	// //start a server
	// if err := http.ListenAndServe(s.addr, s); err != nil {
	// 	log.Fatal(err)
	// }
	api := &api{addr: ":8080"}
	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", api.getUsersHandler)
	mux.HandleFunc("POST /users", api.createUsersHandler)

	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
	// srv := &http.Server{
	// 	Addr:    api.addr,
	// 	Handler: api,
	// }
	// if err := srv.ListenAndServe(); err != nil {
	// 	log.Fatal(err)
	// }

}

//POST
// prajaktabimalkhedkar@Prajaktas-Air go-projects % curl -v -X POST "http://localhost:8080/users" \
// -H "Accept: application/json" \
// -H "Content-Type: application/json" \
// -d '{"first_name": "Prajakta", "last_name": "Bimalkhedkar"}'

// Note: Unnecessary use of -X or --request, POST is already inferred.
// * Host localhost:8080 was resolved.
// * IPv6: ::1
// * IPv4: 127.0.0.1
// *   Trying [::1]:8080...
// * Connected to localhost (::1) port 8080
// > POST /users HTTP/1.1
// > Host: localhost:8080
// > User-Agent: curl/8.7.1
// > Accept: application/json
// > Content-Type: application/json
// > Content-Length: 55
// >
// * upload completely sent off: 55 bytes
// < HTTP/1.1 201 Created
// < Date: Thu, 06 Mar 2025 20:27:45 GMT
// < Content-Length: 0
// <
// * Connection #0 to host localhost left intact

//GET
// prajaktabimalkhedkar@Prajaktas-Air go-projects % curl http://localhost:8080/users

// [{"first_name":"Prajakta","last_name":"Bimalkhedkar"}]
// prajaktabimalkhedkar@Prajaktas-Air go-projects %
